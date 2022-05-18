package product

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type CategoriesQueryParams struct {
	woocommerce.QueryParams
	Search    string   `url:"search,omitempty"`
	Exclude   []string `url:"exclude,omitempty"`
	Include   []string `url:"include,omitempty"`
	HideEmpty bool     `url:"hide_empty,omitempty"`
	Parent    int      `url:"parent,omitempty"`
	Product   int      `url:"product,omitempty"`
	Slug      string   `url:"slug,omitempty"`
}

func (m CategoriesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "name", "slug", "term_group", "description", "count").Error("无效的排序字段"))),
	)
}

func (s service) Categories(params CategoriesQueryParams) (items []entity.ProductCategory, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []entity.ProductCategory
	params.TidyVars()
	urlValues, _ := query.Values(params)
	resp, err := s.woo.Client.R().SetQueryParamsFromValues(urlValues).Get("/products/categories")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			items = res
		}
	}
	return
}

func (s service) Category(id int) (item entity.ProductCategory, err error) {
	var res entity.ProductCategory
	resp, err := s.woo.Client.R().Get(fmt.Sprintf("/products/categories/%d", id))
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

type UpsertCategoryRequest struct {
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Parent      int            `json:"parent"`
	Description string         `json:"description"`
	Display     string               `json:"display,omitempty"`
	Image       *entity.ProductImage `json:"image,omitempty"`
	MenuOrder   int                  `json:"menu_order"`
}

type CreateCategoryRequest = UpsertCategoryRequest
type UpdateCategoryRequest = UpsertCategoryRequest

func (m UpsertCategoryRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name,
			validation.Required.Error("分类名称不能为空"),
		),
	)
}

func (s service) CreateCategory(req CreateCategoryRequest) (item entity.ProductCategory, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.woo.Client.R().SetBody(req).Post("/products/categories")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s service) UpdateCategory(id int, req UpdateCategoryRequest) (item entity.ProductCategory, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.woo.Client.R().SetBody(req).Put(fmt.Sprintf("/products/categories/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s service) DeleteCategory(id int) (item entity.ProductCategory, err error) {
	resp, err := s.woo.Client.R().Delete(fmt.Sprintf("/products/categories/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch tag create,update and delete operation

type CUDCategoriesUpdateRequest struct {
	ID int `json:"id"`
	UpsertTagRequest
}
type CUDCategoriesRequest struct {
	Create []UpsertCategoryRequest      `json:"create"`
	Update []CUDCategoriesUpdateRequest `json:"update"`
	Delete []int                        `json:"delete"`
}

func (m CUDCategoriesRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchCategoriesResult struct {
	Create []entity.ProductTag `json:"create"`
	Update []entity.ProductTag `json:"update"`
	Delete []entity.ProductTag `json:"delete"`
}

func (s service) BatchCategories(req CUDCategoriesRequest) (res BatchCategoriesResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.woo.Client.R().SetBody(req).Post("products/categories/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
