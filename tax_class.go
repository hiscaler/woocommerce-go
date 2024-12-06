package woocommerce

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dashboard-bg/woocommerce-go/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

type taxClassService service

// All List all tax classes
func (s taxClassService) All() (items []entity.TaxClass, err error) {
	resp, err := s.httpClient.R().Get("/tax/classes")
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
		validation.Field(&m.Name, validation.Required.Error("name cannot be empty")),
	)
}

// Create Create a tax class
func (s taxClassService) Create(req CreateTaxClassRequest) (item entity.TaxClass, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/taxes/classes")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete Delete a tax classes
func (s taxClassService) Delete(slug string, force bool) (item entity.TaxClass, err error) {
	slug = strings.TrimSpace(slug)
	if slug == "" {
		err = errors.New("slug parameter cannot be empty")
		return
	}

	resp, err := s.httpClient.R().SetBody(map[string]bool{"force": force}).Delete(fmt.Sprintf("/taxes/classes/%s", slug))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
