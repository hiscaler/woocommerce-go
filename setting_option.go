package woocommerce

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

// https://woocommerce.github.io/woocommerce-rest-api-docs/?php#setting-options

type settingOptionService service

// All list all setting options
func (s settingOptionService) All(settingId int) (items []entity.SettingOption, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/settings/%d", settingId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// One retrieve a setting option
func (s settingOptionService) One(groupId, optionId int) (item entity.SettingOption, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/settings/%d/%d", groupId, optionId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type UpdateSettingOptionRequest struct {
	Value string `json:"value"`
}

func (m UpdateSettingOptionRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Value, validation.Required.Error("设置值不能为空")),
	)
}

func (s settingOptionService) Update(groupId, optionId int, req UpdateSettingOptionRequest) (item entity.SettingOption, err error) {
	if err = req.Validate(); err != nil {
		return
	}
	resp, err := s.httpClient.R().
		SetBody(req).
		Put(fmt.Sprintf("/settings/%d/%d", groupId, optionId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
