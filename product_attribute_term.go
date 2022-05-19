package woocommerce

import (
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productAttributeTermService service

type ProductAttributeTermsQueryParaTermms struct {
	queryParams
}

func (m ProductAttributeTermsQueryParaTermms) Validate() error {
	return nil
}

// All List all product attribute terms
func (s productAttributeTermService) All(attributeId int, params ProductAttributeTermsQueryParaTermms) (items []entity.ProductAttributeTerm, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	urlValues, _ := query.Values(params)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get(fmt.Sprintf("/products/attributes/%d/terms", attributeId))
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
func (s productAttributeTermService) One(attributeId, termId int) (item entity.ProductAttributeTerm, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/attributes/%d/terms/%d", attributeId, termId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateProductAttributeTermRequest struct {
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	MenuOrder   int    `json:"menu_order,omitempty"`
}

func (m CreateProductAttributeTermRequest) Validate() error {
	return nil
}

// Create Create a product attribute term
func (s productAttributeTermService) Create(attributeId int, req CreateProductAttributeTermRequest) (item entity.ProductAttributeTerm, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post(fmt.Sprintf("/products/attributes/%d/terms", attributeId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

type UpdateProductAttributeTermRequest = CreateProductAttributeTermRequest

// Update Update a product attribute term
func (s productAttributeTermService) Update(attributeId, termId int, req UpdateProductAttributeTermRequest) (item entity.ProductAttributeTerm, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/products/attributes/%d/terms/%d", attributeId, termId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete a product attribute term

func (s productAttributeTermService) Delete(attributeId, termId int, force bool) (item entity.ProductAttributeTerm, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/attributes/%d/terms/%d", attributeId, termId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch update product attribute terms

type BatchProductAttributeTermsCreateItem = CreateProductAttributeTermRequest
type BatchProductAttributeTermsUpdateItem struct {
	ID string `json:"id"`
	BatchProductAttributeTermsCreateItem
}

type BatchProductAttributeTermsRequest struct {
	Create []BatchProductAttributeTermsCreateItem `json:"create,omitempty"`
	Update []BatchProductAttributeTermsUpdateItem `json:"update,omitempty"`
	Delete []int                                  `json:"delete,omitempty"`
}

func (m BatchProductAttributeTermsRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchProductAttributeTermsResult struct {
	Create []entity.ProductAttributeTerm `json:"create"`
	Update []entity.ProductAttributeTerm `json:"update"`
	Delete []entity.ProductAttributeTerm `json:"delete"`
}

// Batch Batch create/update/delete product attribute terms
func (s productAttributeTermService) Batch(attributeId int, req BatchProductAttributeTermsRequest) (res BatchProductAttributeTermsResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post(fmt.Sprintf("/products/attributes/%d/batch", attributeId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
