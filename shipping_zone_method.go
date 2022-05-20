package woocommerce

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type shippingZoneMethodService service

// All list all shipping zone methods
func (s shippingZoneMethodService) All(zoneId int) (items []entity.ShippingZoneMethod, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/shipping/zones/%d/methods", zoneId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// One retrieve a shipping zone method
func (s shippingZoneMethodService) One(zoneId, methodId int) (item entity.ShippingZoneMethod, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/shipping/zones/%d/methods/%d", zoneId, methodId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Include include a shipping method to a shipping zone

type ShippingZoneMethodIncludeRequest struct {
	MethodId string `json:"method_id"`
}

func (m ShippingZoneMethodIncludeRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.MethodId, validation.Required.Error("配送方式不能为空")),
	)
}

func (s shippingZoneMethodService) Include(zoneId int, req ShippingZoneMethodIncludeRequest) (item entity.ShippingZoneMethod, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
		SetBody(req).
		Post(fmt.Sprintf("/shipping/zones/%d/methods", zoneId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Update

type UpdateShippingZoneMethodSetting struct {
	Value string `json:"value"`
}

type UpdateShippingZoneMethodRequest struct {
	Order    int                             `json:"order"`
	Enabled  bool                            `json:"enabled"`
	Settings UpdateShippingZoneMethodSetting `json:"settings"`
}

func (m UpdateShippingZoneMethodRequest) Validate() error {
	return nil
}

func (s shippingZoneMethodService) Update(zoneId, methodId int, req UpdateShippingZoneMethodRequest) (item entity.ShippingZoneMethod, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
		SetBody(req).
		Put(fmt.Sprintf("/shipping/zones/%d/methods/%d", zoneId, methodId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete delete a shipping zone
func (s shippingZoneMethodService) Delete(zoneId, methodId int, force bool) (item entity.ShippingZone, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/shipping/zones/%d/%d", zoneId, methodId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
