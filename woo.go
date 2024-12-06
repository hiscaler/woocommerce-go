// Package woocommerce is a Woo Commerce lib.
//
// Quick start:
//
//	b, err := os.ReadFile("./config/config_test.json")
//	if err != nil {
//	   panic(fmt.Sprintf("Read config error: %s", err.Error()))
//	}
//	var c config.Config
//	err = jsoniter.Unmarshal(b, &c)
//	if err != nil {
//	   panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
//	}
//
//	wooClient = NewClient(c)
//	// Query an order
//	order, err := wooClient.Services.Order.One(1)
//	if err != nil {
//	   fmt.Println(err)
//	} else {
//	   fmt.Println(fmt.Sprintf("%#v", order))
//	}
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
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/dashboard-bg/woocommerce-go/config"
	"github.com/go-resty/resty/v2"
	"github.com/hiscaler/gox/inx"
	"github.com/hiscaler/gox/stringx"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

const (
	Version       = "1.0.3"
	UserAgent     = "WooCommerce API Client-Golang/" + Version
	HashAlgorithm = "HMAC-SHA256"
)

// https://woocommerce.github.io/woocommerce-rest-api-docs/?php#request-response-format
const (
	BadRequestError           = 400
	UnauthorizedError         = 401
	NotFoundError             = 404
	InternalServerError       = 500
	MethodNotImplementedError = 501
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

// OAuth signature
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

// NewClient Creates a new WooCommerce client
//
// You must give a config with NewClient method params.
// After you can operate data use this client.
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
			for k, vs := range request.QueryParam {
				var v string
				switch len(vs) {
				case 0:
					continue
				case 1:
					v = vs[0]
				default:
					// if is array params, must convert to string, example: status=1&status=2 replace to status=1,2
					v = strings.Join(vs, ",")
				}
				params.Set(k, v)
			}
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

// ErrorWrap wrap an error, if status code is 200, return nil, otherwise return an error
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
			message = "Bad request"
		case UnauthorizedError:
			message = "Unauthorized operation, please confirm whether you have permission"
		case NotFoundError:
			message = "Resource not found"
		case InternalServerError:
			message = "Server internal error"
		case MethodNotImplementedError:
			message = "method not implemented"
		default:
			message = "Unknown error"
		}
	}
	return fmt.Errorf("%d: %s", code, message)
}
