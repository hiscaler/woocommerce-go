package woocommerce

import (
	"fmt"

	"github.com/dashboard-bg/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type paymentGatewayService service

func (s paymentGatewayService) All() (items []entity.PaymentGateway, err error) {
	resp, err := s.httpClient.R().Get("/payment_gateways")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// One retrieve a payment gateway
func (s paymentGatewayService) One(id string) (item entity.PaymentGateway, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/payment_gateways/%s", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Update

type UpdatePaymentGatewayRequest struct {
	Title             string                                  `json:"title,omitempty"`
	Description       string                                  `json:"description,omitempty"`
	Order             int                                     `json:"order,omitempty"`
	Enabled           bool                                    `json:"enabled,omitempty"`
	MethodTitle       string                                  `json:"method_title,omitempty"`
	MethodDescription string                                  `json:"method_description,omitempty"`
	MethodSupports    []string                                `json:"method_supports,omitempty"`
	Settings          map[string]entity.PaymentGatewaySetting `json:"settings,omitempty"`
}

func (m UpdatePaymentGatewayRequest) Validate() error {
	return nil
}

func (s paymentGatewayService) Update(id string, req UpdatePaymentGatewayRequest) (item entity.PaymentGateway, err error) {
	if err = req.Validate(); err != nil {
		return
	}
	resp, err := s.httpClient.R().
		SetBody(req).
		Put(fmt.Sprintf("/payment_gateways/%s", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
