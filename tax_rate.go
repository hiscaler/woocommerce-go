package woocommerce

import (
	"errors"
	"fmt"

	"github.com/dashboard-bg/woocommerce-go/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

type taxRateService service

type TaxRatesQueryParams struct {
	queryParams
	Class string `url:"class,omitempty"`
}

func (m TaxRatesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "order", "priority").Error("invalid sort field"))),
	)
}

// All List all tax rate
func (s taxRateService) All(params TaxRatesQueryParams) (items []entity.TaxRate, total, totalPages int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	resp, err := s.httpClient.R().SetQueryParamsFromValues(toValues(params)).Get("/taxes")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
		total, totalPages, isLastPage = parseResponseTotal(params.Page, resp)
	}
	return
}

// One Retrieve a tax rate
func (s taxRateService) One(id int) (item entity.TaxRate, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/taxes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateTaxRateRequest struct {
	Country   string   `json:"country,omitempty"`
	State     string   `json:"state,omitempty"`
	Postcode  string   `json:"postcode,omitempty"`
	City      string   `json:"city,omitempty"`
	Postcodes []string `json:"postcodes,omitempty"`
	Cities    []string `json:"cities,omitempty"`
	Rate      string   `json:"rate,omitempty"`
	Name      string   `json:"name,omitempty"`
	Priority  int      `json:"priority,omitempty"`
	Compound  bool     `json:"compound,omitempty"`
	Shipping  bool     `json:"shipping,omitempty"`
	Order     int      `json:"order,omitempty"`
	Class     string   `json:"class,omitempty"`
}

func (m CreateTaxRateRequest) Validate() error {
	return nil
}

// Create Create a product attribute
func (s taxRateService) Create(req CreateTaxRateRequest) (item entity.TaxRate, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/taxes")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

type UpdateTaxRateRequest = CreateTaxRateRequest

// Update Update a tax rate
func (s taxRateService) Update(id int, req UpdateTaxRateRequest) (item entity.TaxRate, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/taxes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete a tax rate
func (s taxRateService) Delete(id int, force bool) (item entity.TaxRate, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/taxes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch update tax rates

type BatchTaxRatesCreateItem = CreateTaxRateRequest
type BatchTaxRatesUpdateItem struct {
	ID string `json:"id"`
	BatchTaxRatesCreateItem
}

type BatchTaxRatesRequest struct {
	Create []BatchTaxRatesCreateItem `json:"create,omitempty"`
	Update []BatchTaxRatesUpdateItem `json:"update,omitempty"`
	Delete []int                     `json:"delete,omitempty"`
}

func (m BatchTaxRatesRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("invalid request data")
	}
	return nil
}

type BatchTaxRatesResult struct {
	Create []entity.TaxRate `json:"create"`
	Update []entity.TaxRate `json:"update"`
	Delete []entity.TaxRate `json:"delete"`
}

// Batch Batch create/update/delete tax rates
func (s taxRateService) Batch(req BatchTaxRatesRequest) (res BatchTaxRatesResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/taxes/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
