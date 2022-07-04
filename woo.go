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
	"github.com/hiscaler/gox/stringx"
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
	"unsafe"
)

const (
	Version       = "1.0.3"
	UserAgent     = "WooCommerce API Client-Golang/" + Version
	HashAlgorithm = "HMAC-SHA256"
)

// https://woocommerce.github.io/woocommerce-rest-api-docs/?php#request-response-format
const (
	BadRequestError         = 400 // 错误的请求
	UnauthorizedError       = 401 // 身份验证或权限错误
	NotFoundError           = 404 // 访问资源不存在
	InternalServerError     = 500 // 服务器内部错误
	MethodNotImplementedErr = 501 // 方法未实现
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
	Coupon               couponService
	Customer             customerService
	Order                orderService
	OrderNote            orderNoteService
	OrderRefund          orderRefundService
	Product              productService
	ProductVariation     productVariationService
	ProductAttribute     productAttributeService
	ProductAttributeTerm productAttributeTermService
	ProductCategory      productCategoryService
	ProductShippingClass productShippingClassService
	ProductTag           productTagService
	ProductReview        productReviewService
	Report               reportService
	TaxRate              taxRateService
	TaxClass             taxClassService
	Webhook              webhookService
	Setting              settingService
	SettingOption        settingOptionService
	PaymentGateway       paymentGatewayService
	ShippingZone         shippingZoneService
	ShippingZoneLocation shippingZoneLocationService
	ShippingZoneMethod   shippingZoneMethodService
	ShippingMethod       shippingMethodService
	SystemStatus         systemStatusService
	SystemStatusTool     systemStatusToolService
	Data                 dataService
}

// OAuth 签名
func oauthSignature(config config.Config, method, endpoint string, params url.Values) string {
	sb := strings.Builder{}
	sb.WriteString(config.ConsumerSecret)
	if config.Version != "v1" && config.Version != "v2" {
		sb.WriteByte('&')
	}
	consumerSecret := sb.String()

	sb.Reset()
	sb.WriteString(method)
	sb.WriteByte('&')
	sb.WriteString(url.QueryEscape(endpoint))
	sb.WriteByte('&')
	sb.WriteString(url.QueryEscape(params.Encode()))
	mac := hmac.New(sha256.New, stringx.ToBytes(consumerSecret))
	mac.Write(stringx.ToBytes(sb.String()))
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
		if !inx.StringIn(config.Version, "v1", "v2", "v3") {
			config.Version = "v3"
		}
	}
	if config.Timeout < 2 {
		config.Timeout = 2
	}

	httpClient := resty.New().
		SetDebug(config.Debug).
		SetBaseURL(strings.TrimRight(config.URL, "/") + "/wp-json/wc/" + config.Version).
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
				if config.AddAuthenticationToURL {
					params.Add("consumer_key", config.ConsumerKey)
					params.Add("consumer_secret", config.ConsumerSecret)
				} else {
					// Set to header
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

	jsoniter.RegisterTypeDecoderFunc("float64", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		switch iter.WhatIsNext() {
		case jsoniter.StringValue:
			var t float64
			v := strings.TrimSpace(iter.ReadString())
			if v != "" {
				var err error
				if t, err = strconv.ParseFloat(v, 64); err != nil {
					iter.Error = err
					return
				}
			}
			*((*float64)(ptr)) = t
		default:
			*((*float64)(ptr)) = iter.ReadFloat64()
		}
	})
	httpClient.JSONMarshal = jsoniter.Marshal
	httpClient.JSONUnmarshal = jsoniter.Unmarshal
	xService := service{
		debug:      config.Debug,
		logger:     logger,
		httpClient: httpClient,
	}
	wooClient.Services = services{
		Coupon:               (couponService)(xService),
		Customer:             (customerService)(xService),
		Order:                (orderService)(xService),
		OrderNote:            (orderNoteService)(xService),
		OrderRefund:          (orderRefundService)(xService),
		Product:              (productService)(xService),
		ProductVariation:     (productVariationService)(xService),
		ProductAttribute:     (productAttributeService)(xService),
		ProductAttributeTerm: (productAttributeTermService)(xService),
		ProductCategory:      (productCategoryService)(xService),
		ProductShippingClass: (productShippingClassService)(xService),
		ProductTag:           (productTagService)(xService),
		ProductReview:        (productReviewService)(xService),
		Report:               (reportService)(xService),
		TaxRate:              (taxRateService)(xService),
		TaxClass:             (taxClassService)(xService),
		Webhook:              (webhookService)(xService),
		Setting:              (settingService)(xService),
		SettingOption:        (settingOptionService)(xService),
		PaymentGateway:       (paymentGatewayService)(xService),
		ShippingZone:         (shippingZoneService)(xService),
		ShippingZoneLocation: (shippingZoneLocationService)(xService),
		ShippingZoneMethod:   (shippingZoneMethodService)(xService),
		ShippingMethod:       (shippingMethodService)(xService),
		SystemStatus:         (systemStatusService)(xService),
		SystemStatusTool:     (systemStatusToolService)(xService),
		Data:                 (dataService)(xService),
	}
	return wooClient
}

// Parse response header, get total and total pages, and check it is last page
func parseResponseTotal(currentPage int, resp *resty.Response) (total, totalPages int, isLastPage bool) {
	if currentPage == 0 {
		currentPage = 1
	}
	value := resp.Header().Get("X-Wp-Total")
	if value != "" {
		total, _ = strconv.Atoi(value)
	}

	value = resp.Header().Get("X-Wp-Totalpages")
	if value != "" {
		totalPages, _ = strconv.Atoi(value)
	}
	isLastPage = currentPage >= totalPages
	return
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
		case MethodNotImplementedErr:
			message = "方法未实现"
		default:
			message = "未知错误"
		}
	}
	return fmt.Errorf("%d: %s", code, message)
}
