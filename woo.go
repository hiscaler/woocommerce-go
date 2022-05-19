package woocommerce

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/inx"
	"github.com/hiscaler/woocommerce-go/config"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Version       = "1.0.0"
	UserAgent     = "WooCommerce API Client-Golang/" + Version
	HashAlgorithm = "HMAC-SHA256"
)

// https://woocommerce.github.io/woocommerce-rest-api-docs/?php#request-response-format
const (
	BadRequestError     = 400 // 错误的请求
	UnauthorizedError   = 401 // 身份验证或权限错误
	NotFoundError       = 404 // 访问资源不存在
	InternalServerError = 500 // 服务器内部错误
)

var ErrNotFound = errors.New("WooCommerce: not found")

func init() {
	extra.RegisterFuzzyDecoders()
}

type WooCommerce struct {
	Debug    bool        // Is debug mode
	Logger   *log.Logger // Log
	Services services    // WooCommerce API services
}

type service struct {
	debug      bool          // Is debug mode
	logger     *log.Logger   // Log
	httpClient *resty.Client // HTTP client
}

type services struct {
	Product          productService
	ProductTag       productTagService
	ProductCategory  productCategoryService
	ProductVariation productVariationService
	Customer         customerService
	Order            orderService
	OrderNote        orderNoteService
	OrderRefund      orderRefundService
	Setting          settingService
	SettingOption    settingOptionService
	TaxClass         taxClassService
	System           systemService
	Coupon           couponService
	TaxRate          taxRateService
	PaymentGateway   paymentGatewayService
}

// OAuth 签名
func oauthSignature(config config.Config, method, endpoint string, params url.Values) string {
	consumerSecret := config.ConsumerSecret
	if config.Version != "v1" && config.Version != "v2" {
		consumerSecret = consumerSecret + "&"
	}
	s := method + "&" + url.QueryEscape(endpoint) + "&" + url.QueryEscape(params.Encode())
	mac := hmac.New(sha256.New, []byte(consumerSecret))
	mac.Write([]byte(s))
	signatureBytes := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signatureBytes)
}

func NewClient(config config.Config) *WooCommerce {
	logger := log.New(os.Stdout, "[ WooCommerce ] ", log.LstdFlags|log.Llongfile)
	wooClient := &WooCommerce{
		Debug:  config.Debug,
		Logger: logger,
	}
	// Add default value
	if config.Version == "" {
		config.Version = "v3"
	} else {
		config.Version = strings.ToLower(config.Version)
		if !inx.StringIn(config.Version, "v1", "v1", "v3") {
			config.Version = "v3"
		}
	}
	if config.Timeout < 2 {
		config.Timeout = 2
	}

	storeURL := strings.Trim(config.URL, "/") + "/wp-json/wc/" + config.Version
	httpClient := resty.New().
		SetDebug(config.Debug).
		SetBaseURL(storeURL).
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
			"User-Agent":   UserAgent,
		}).
		SetAllowGetMethodPayload(true).
		SetTimeout(config.Timeout * time.Second).
		SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.VerifySSL},
			DialContext: (&net.Dialer{
				Timeout: config.Timeout * time.Second,
			}).DialContext,
		}).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			params := url.Values{}
			if strings.HasPrefix(config.URL, "https") {
				// basicAuth
				if config.UseAuthInQuery {
					params.Add("consumer_key", config.ConsumerKey)
					params.Add("consumer_secret", config.ConsumerSecret)
				} else {
					client.SetAuthScheme("Basic").
						SetAuthToken(fmt.Sprintf("%s %s", config.ConsumerKey, config.ConsumerSecret))
				}
			} else {
				// oAuth
				params.Add("oauth_consumer_key", config.ConsumerKey)
				params.Add("oauth_timestamp", strconv.Itoa(int(time.Now().Unix())))
				nonce := make([]byte, 16)
				rand.Read(nonce)
				sha1Nonce := fmt.Sprintf("%x", sha1.Sum(nonce))
				params.Add("oauth_nonce", sha1Nonce)
				params.Add("oauth_signature_method", HashAlgorithm)
				for k, vs := range request.QueryParam {
					for _, v := range vs {
						params.Add(k, v)
					}
				}
				params.Add("oauth_signature", oauthSignature(config, request.Method, client.BaseURL+request.URL, params))
			}
			request.QueryParam = params
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, response *resty.Response) (err error) {
			if response.IsError() {
				r := struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				}{}
				if err = jsoniter.Unmarshal(response.Body(), &r); err == nil {
					err = ErrorWrap(response.StatusCode(), r.Message)
				}
			}
			if err != nil {
				logger.Printf("OnAfterResponse error: %s", err.Error())
			}
			return
		})
	if config.Debug {
		httpClient.EnableTrace()
	}
	httpClient.JSONMarshal = jsoniter.Marshal
	httpClient.JSONUnmarshal = jsoniter.Unmarshal
	xService := service{
		debug:      config.Debug,
		logger:     logger,
		httpClient: httpClient,
	}
	wooClient.Services = services{
		Product:          (productService)(xService),
		ProductTag:       (productTagService)(xService),
		ProductCategory:  (productCategoryService)(xService),
		ProductVariation: (productVariationService)(xService),
		Customer:         (customerService)(xService),
		Order:            (orderService)(xService),
		OrderNote:        (orderNoteService)(xService),
		OrderRefund:      (orderRefundService)(xService),
		Setting:          (settingService)(xService),
		SettingOption:    (settingOptionService)(xService),
		TaxClass:         (taxClassService)(xService),
		System:           (systemService)(xService),
		Coupon:           (couponService)(xService),
		TaxRate:          (taxRateService)(xService),
		PaymentGateway:   (paymentGatewayService)(xService),
	}
	return wooClient
}

// ErrorWrap 错误包装
func ErrorWrap(code int, message string) error {
	if code == http.StatusOK {
		return nil
	}

	if code == NotFoundError {
		return ErrNotFound
	}

	message = strings.TrimSpace(message)
	if message == "" {
		switch code {
		case BadRequestError:
			message = "错误的请求"
		case UnauthorizedError:
			message = "身份验证或权限错误"
		case NotFoundError:
			message = "访问资源不存在"
		case InternalServerError:
			message = "服务器内部错误"
		default:
			message = "未知错误"
		}
	}
	return fmt.Errorf("%d: %s", code, message)
}
