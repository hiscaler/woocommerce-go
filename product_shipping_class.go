package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productShippingClassService service

type ProductShippingClassesQueryParams struct {
	queryParams
	Search    string   `url:"search,omitempty"`
	Exclude   []string `url:"exclude,omitempty"`
	Include   []string `url:"include,omitempty"`
	HideEmpty bool     `url:"hide_empty,omitempty"`
	Parent    int      `url:"parent,omitempty"`
	Product   int      `url:"product,omitempty"`
	Slug      string   `url:"slug,omitempty"`
}

func (m ProductShippingClassesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "name", "slug", "term_group", "description", "count").Error("无效的排序类型"))),
	)
}

// All List all product shipping class
func (s productShippingClassService) All(params ProductShippingClassesQueryParams) (items []entity.ProductShippingClass, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	urlValues, _ := query.Values(params)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get("/products/shipping_classes")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &items); err == nil {
			isLastPage = len(items) < params.PerPage
		}
	}
	return
}

// One Retrieve a product shipping class
func (s productShippingClassService) One(id int) (item entity.ProductShippingClass, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/shipping_classes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateProductShippingClassRequest struct {
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}

func (m CreateProductShippingClassRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required.Error("名称不能为空")),
	)
}

// Create Create a product shipping class
func (s productShippingClassService) Create(req CreateProductShippingClassRequest) (item entity.ProductShippingClass, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/shipping_classes")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

type UpdateProductShippingClassRequest = CreateProductShippingClassRequest

// Update Update a product shipping class
func (s productShippingClassService) Update(id int, req UpdateProductShippingClassRequest) (item entity.ProductShippingClass, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/products/shipping_classes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete a product shipping class
func (s productShippingClassService) Delete(id int, force bool) (item entity.ProductShippingClass, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/shipping_classes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch update product shipping classes

type BatchProductShippingClassesCreateItem = CreateProductShippingClassRequest
type BatchProductShippingClassesUpdateItem struct {
	ID string `json:"id"`
	BatchProductShippingClassesCreateItem
}

type BatchProductShippingClassesRequest struct {
	Create []BatchProductShippingClassesCreateItem `json:"create,omitempty"`
	Update []BatchProductShippingClassesUpdateItem `json:"update,omitempty"`
	Delete []int                                   `json:"delete,omitempty"`
}

func (m BatchProductShippingClassesRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchProductShippingClassesResult struct {
	Create []entity.ProductShippingClass `json:"create"`
	Update []entity.ProductShippingClass `json:"update"`
	Delete []entity.ProductShippingClass `json:"delete"`
}

// Batch Batch create/update/delete product shipping classes
func (s productShippingClassService) Batch(req BatchProductShippingClassesRequest) (res BatchProductShippingClassesResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/shipping_classes/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
