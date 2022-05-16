package product

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity/product"
	jsoniter "github.com/json-iterator/go"
)

type CategoriesQueryParams struct {
	Context   string   `json:"context,omitempty"`
	Search    string   `json:"search,omitempty"`
	Exclude   []string `json:"exclude,omitempty"`
	Include   []string `json:"include,omitempty"`
	Order     string   `json:"order,omitempty"`
	OrderBy   string   `json:"orderby,omitempty"`
	HideEmpty bool     `json:"hide_empty,omitempty"`
	Parent    int      `json:"parent"`
	Product   int      `json:"product,omitempty"`
	Slug      string   `json:"slug,omitempty"`
}

func (m CategoriesQueryParams) Validate() error {
	return nil
}

func (s service) Categories(params CategoriesQueryParams) (items []product.Category, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []product.Category
	qp := make(map[string]string, 0)
	resp, err := s.woo.Client.R().
		SetQueryParams(qp).
		Get("/products/categories")
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

func (s service) Category(id int) (item product.Category, err error) {
	var res product.Category
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
	Display     string         `json:"display,omitempty"`
	Image       *product.Image `json:"image,omitempty"`
	MenuOrder   int            `json:"menu_order"`
}

func (m UpsertCategoryRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name,
			validation.Required.Error("分类名称不能为空"),
		),
	)
}

func (s service) CreateCategory(req UpsertCategoryRequest) (item product.Category, err error) {
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

func (s service) UpdateCategory(id int, req UpsertCategoryRequest) (item product.Category, err error) {
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

func (s service) DeleteCategory(id int) (item product.Category, err error) {
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
	Create []product.Tag `json:"create"`
	Update []product.Tag `json:"update"`
	Delete []product.Tag `json:"delete"`
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
