package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productAttributeService service

type ProductAttributesQueryParams struct {
	queryParams
}

func (m ProductAttributesQueryParams) Validate() error {
	return nil
}

// All List all product attributes
func (s productAttributeService) All(params ProductAttributesQueryParams) (items []entity.ProductAttribute, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	urlValues, _ := query.Values(params)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get("/products/attributes")
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

// One Retrieve a product attribute
func (s productAttributeService) One(id int) (item entity.ProductAttribute, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/attributes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateProductAttributeRequest struct {
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Type        string `json:"type,omitempty"`
	OrderBy     string `json:"order_by,omitempty"`
	HasArchives bool   `json:"has_archives,omitempty"`
}

func (m CreateProductAttributeRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("menu_order", "name", "name_num", "id").Error("无效的排序方式"))),
	)
}

// Create Create a product attribute
func (s productAttributeService) Create(req CreateProductAttributeRequest) (item entity.ProductAttribute, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/attributes")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

type UpdateProductAttributeRequest = CreateProductAttributeRequest

// Update Update a product attribute
func (s productAttributeService) Update(id int, req UpdateProductAttributeRequest) (item entity.ProductAttribute, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/products/attributes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete a product attribute

func (s productAttributeService) Delete(id int, force bool) (item entity.ProductAttribute, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/attributes/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch update porduct attributes

type BatchProductAttributesCreateItem = CreateProductAttributeRequest
type BatchProductAttributesUpdateItem struct {
	ID string `json:"id"`
	BatchProductAttributesCreateItem
}

type BatchProductAttributesRequest struct {
	Create []BatchProductAttributesCreateItem `json:"create,omitempty"`
	Update []BatchProductAttributesUpdateItem `json:"update,omitempty"`
	Delete []int                              `json:"delete,omitempty"`
}

func (m BatchProductAttributesRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchProductAttributesResult struct {
	Create []entity.ProductAttribute `json:"create"`
	Update []entity.ProductAttribute `json:"update"`
	Delete []entity.ProductAttribute `json:"delete"`
}

// Batch Batch create/update/delete product attributes
func (s productAttributeService) Batch(req BatchProductAttributesRequest) (res BatchProductAttributesResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/attributes/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
