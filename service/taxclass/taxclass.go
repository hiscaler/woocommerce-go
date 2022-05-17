package taxclass

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

type TaxClass struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// TaxClasses List all tax classes
func (s service) TaxClasses() (items []TaxClass, err error) {
	resp, err := s.woo.Client.R().Get("/tax/classes")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// Create tax class request

type CreateTaxClassRequest struct {
	Name string `json:"name"`
}

func (m CreateTaxClassRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required.Error("名称不能为空")),
	)
}

func (s service) CreateTaxClass(req CreateTaxClassRequest) (item TaxClass, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.woo.Client.R().SetBody(req).Post("/taxes/classes")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// DeleteTaxClass Delete tax classes
func (s service) DeleteTaxClass(slug string) (item TaxClass, err error) {
	slug = strings.TrimSpace(slug)
	if slug == "" {
		err = errors.New("slug 参数不能为空")
		return
	}

	resp, err := s.woo.Client.R().Delete(fmt.Sprintf("/taxes/classes/%s", slug))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
