package woocommerce

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/woocommerce-go/config"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Version       = "1.0.0"
	UserAgent     = "WooCommerce API Client-Golang/" + Version
	HashAlgorithm = "HMAC-SHA256"
)

var ErrNotFound = errors.New("WooCommerce: not found")

type queryDefaultValues struct {
	Page     int `json:"page"`     // 当前页
	PageSize int `json:"per_page"` // 每页数据量
}

type WooCommerce struct {
	Debug              bool               // 是否调试模式
	Client             *resty.Client      // HTTP 客户端
	Logger             *log.Logger        // 日志
	QueryDefaultValues queryDefaultValues // 查询默认值
}

// OAuth 签名
func oauthSignature(config config.Config, method, endpoint, params string) string {
	signingKey := config.ConsumerKey
	if config.Version != "v1" && config.Version != "v2" {
		signingKey = signingKey + "&"
	}

	a := strings.Join([]string{method, url.QueryEscape(endpoint), url.QueryEscape(params)}, "&")
	mac := hmac.New(sha256.New, []byte(signingKey))
	mac.Write([]byte(a))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes)
}

func NewClient(config config.Config) *WooCommerce {
	logger := log.New(os.Stdout, "[ ShipOut ] ", log.LstdFlags|log.Llongfile)
	wooInstance := &WooCommerce{
		Debug:  config.Debug,
		Logger: logger,
		QueryDefaultValues: queryDefaultValues{
			Page:     1,
			PageSize: 100,
		},
	}
	// Add default value
	if config.Version == "" {
		config.Version = "v3"
	}

	storeURL := config.URL + "/wp-json/wc/" + config.Version
	client := resty.New().
		SetDebug(config.Debug).
		SetBaseURL(config.URL).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   UserAgent,
		}).
		SetAllowGetMethodPayload(true).
		SetTimeout(10 * time.Second).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			params := url.Values{}
			if strings.HasPrefix(storeURL, "https") {
				// basicAuth
				params.Add("consumer_key", config.ConsumerKey)
				params.Add("consumer_secret", config.ConsumerSecret)
			} else {
				// oAuth
				params.Add("oauth_consumer_key", config.ConsumerKey)
				params.Add("oauth_timestamp", strconv.Itoa(int(time.Now().Unix())))
				nonce := make([]byte, 16)
				rand.Read(nonce)
				sha1Nonce := fmt.Sprintf("%x", sha1.Sum(nonce))
				params.Add("oauth_nonce", sha1Nonce)
				params.Add("oauth_signature_method", HashAlgorithm)
				var keys []string
				for k, _ := range params {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				var paramStrs []string
				for _, key := range keys {
					paramStrs = append(paramStrs, fmt.Sprintf("%s=%s", key, params.Get(key)))
				}
				paramStr := strings.Join(paramStrs, "&")
				params.Add("oauth_signature", oauthSignature(config, request.Method, request.URL, paramStr))
			}
			request.SetQueryParamsFromValues(params)
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			if response.IsSuccess() {
				r := struct {
					Result    string `json:"result"`
					ErrorCode string `json:"ErrorCode"`
					Message   string `json:"message"`
				}{}
				if err = jsoniter.Unmarshal(response.Body(), &r); err == nil {
					code := r.ErrorCode
					if code == "" {
						code = r.Result
					}
					err = ErrorWrap(code, r.Message)
				}
			}
			if err != nil {
				logger.Printf("OnAfterResponse error: %s", err.Error())
			}
			return
		})
	if config.Debug {
		client.EnableTrace()
	}
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	wooInstance.Client = client
	return wooInstance
}

// ErrorWrap 错误包装
func ErrorWrap(code string, message string) error {
	if code == "" || code == "OK" {
		return nil
	}

	message = strings.TrimSpace(message)
	if message == "" {
	}
	return fmt.Errorf("%s: %s", code, message)
}
