package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productCategoryService service

type ProductCategoriesQueryParams struct {
	queryParams
	Search    string `url:"search,omitempty"`
	Exclude   []int  `url:"exclude,omitempty"`
	Include   []int  `url:"include,omitempty"`
	HideEmpty bool   `url:"hide_empty,omitempty"`
	Parent    int    `url:"parent,omitempty"`
	Product   int    `url:"product,omitempty"`
	Slug      string `url:"slug,omitempty"`
}

func (m ProductCategoriesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "name", "slug", "term_group", "description", "count").Error("无效的排序字段"))),
	)
}

func (s productCategoryService) All(params ProductCategoriesQueryParams) (items []entity.ProductCategory, total, totalPages int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	resp, err := s.httpClient.R().SetQueryParamsFromValues(toValues(params)).Get("/products/categories")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
		total, totalPages, isLastPage = parseResponseTotal(params.Page, resp)
	}
	return
}

func (s productCategoryService) One(id int) (item entity.ProductCategory, err error) {
	var res entity.ProductCategory
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/categories/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			item = res
		}
	}
	return
}

// 新增商品标签

type UpsertProductCategoryRequest struct {
	Name        string               `json:"name"`
	Slug        string               `json:"slug,omitempty"`
	Parent      int                  `json:"parent,omitempty"`
	Description string               `json:"description,omitempty"`
	Display     string               `json:"display,omitempty"`
	Image       *entity.ProductImage `json:"image,omitempty"`
	MenuOrder   int                  `json:"menu_order,omitempty"`
}

type CreateProductCategoryRequest = UpsertProductCategoryRequest
type UpdateProductCategoryRequest = UpsertProductCategoryRequest

func (m UpsertProductCategoryRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name,
			validation.Required.Error("分类名称不能为空"),
		),
	)
}

func (s productCategoryService) Create(req CreateProductCategoryRequest) (item entity.ProductCategory, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/categories")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s productCategoryService) Update(id int, req UpdateProductCategoryRequest) (item entity.ProductCategory, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/products/categories/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s productCategoryService) Delete(id int, force bool) (item entity.ProductCategory, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/categories/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch category create,update and delete operation

type BatchProductCategoriesCreateItem = UpsertProductCategoryRequest
type BatchProductCategoriesUpdateItem struct {
	ID int `json:"id"`
	UpsertProductTagRequest
}
type BatchProductCategoriesRequest struct {
	Create []BatchProductCategoriesCreateItem `json:"create,omitempty"`
	Update []BatchProductCategoriesUpdateItem `json:"update,omitempty"`
	Delete []int                              `json:"delete,omitempty"`
}

func (m BatchProductCategoriesRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchProductCategoriesResult struct {
	Create []entity.ProductTag `json:"create"`
	Update []entity.ProductTag `json:"update"`
	Delete []entity.ProductTag `json:"delete"`
}

func (s productCategoryService) Batch(req BatchProductCategoriesRequest) (res BatchProductCategoriesResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/categories/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
