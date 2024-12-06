package woocommerce

import (
	"fmt"

	"github.com/dashboard-bg/woocommerce-go/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

type shippingZoneService service

// All list all shipping zones
func (s shippingZoneService) All() (items []entity.ShippingZone, err error) {
	resp, err := s.httpClient.R().Get("/shipping/zones")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// One retrieve a shipping zone
func (s shippingZoneService) One(id int) (item entity.ShippingZone, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/shipping/zones/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateShippingZoneRequest struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func (m CreateShippingZoneRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required.Error("name cannot be empty")),
		validation.Field(&m.Order, validation.Min(0).Error("sort value cannot be less than {{.threshold}}")),
	)
}

func (s shippingZoneService) Create(req CreateShippingZoneRequest) (item entity.ShippingZone, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/shipping/zones")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Update

type UpdateShippingZoneRequest = CreateShippingZoneRequest

func (s shippingZoneService) Update(id int, req UpdateShippingZoneRequest) (item entity.ShippingZone, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/shipping/zones/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete delete a shipping zone
func (s shippingZoneService) Delete(id int, force bool) (item entity.ShippingZone, err error) {
	resp, err := s.httpClient.R().SetBody(map[string]bool{"force": force}).Delete(fmt.Sprintf("/shipping/zones/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
