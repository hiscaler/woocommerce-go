package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type shippingZoneLocationService service

// All list all shipping zone locations
func (s shippingZoneLocationService) All(shippingZoneId int) (items []entity.ShippingZoneLocation, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/shipping/zones/%d/locations", shippingZoneId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// Update

type UpdateShippingZoneLocationsRequest []entity.ShippingZoneLocation

func (m UpdateShippingZoneLocationsRequest) Validate() error {
	return validation.Validate(m, validation.Required.Error("待更新数据不能为空"),
		validation.By(func(value interface{}) error {
			items, ok := value.([]entity.ShippingZoneLocation)
			if !ok {
				return errors.New("待更新数据错误")
			}

			for _, item := range items {
				err := validation.ValidateStruct(&item,
					validation.Field(&item.Code, validation.Required.Error("代码不能为空")),
					validation.Field(&item.Type, validation.In("postcode", "country", "state", "continent").Error("类型错误")),
				)
				if err != nil {
					return err
				}
			}
			return nil
		}),
	)
}

func (s shippingZoneLocationService) Update(shippingZoneId int, req UpdateShippingZoneLocationsRequest) (items []entity.ShippingZoneLocation, err error) {
	if err = req.Validate(); err != nil {
		return
	}
	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/shipping/zones/%d/locations", shippingZoneId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}
