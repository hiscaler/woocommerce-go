package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity/product"
	jsoniter "github.com/json-iterator/go"
)

type productTagService service

type ProductTagsQueryParams struct {
	QueryParams
	Search    string   `url:"search,omitempty"`
	Exclude   []string `url:"exclude,omitempty"`
	Include   []string `url:"include,omitempty"`
	HideEmpty bool     `url:"hide_empty,omitempty"`
	Product   int      `url:"product,omitempty"`
	Slug      string   `url:"slug,omitempty"`
}

func (m ProductTagsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "name", "slug", "term_group", "description", "count").Error("无效的排序字段"))),
	)
}

func (s productTagService) All(params ProductTagsQueryParams) (items []product.Tag, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []product.Tag
	params.TidyVars()
	urlValues, _ := query.Values(params)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get("/products/tags")
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

func (s productTagService) One(id int) (item product.Tag, err error) {
	var res product.Tag
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/tags/%d", id))
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

type UpsertTagRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateTagRequest = UpsertTagRequest
type UpdateTagRequest = UpsertTagRequest

func (m UpsertTagRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name,
			validation.Required.Error("标签名称不能为空"),
		),
	)
}

func (s productTagService) Create(req CreateTagRequest) (item product.Tag, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/tags")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s productTagService) Update(id int, req UpdateTagRequest) (item product.Tag, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/products/tags/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s productTagService) Delete(id int) (item product.Tag, err error) {
	resp, err := s.httpClient.R().Delete(fmt.Sprintf("/products/tags/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch tag create,update and delete operation

type CUDTagsUpdateRequest struct {
	ID int `json:"id"`
	UpsertTagRequest
}
type CUDTagsRequest struct {
	Create []UpsertTagRequest     `json:"create"`
	Update []CUDTagsUpdateRequest `json:"update"`
	Delete []int                  `json:"delete"`
}

func (m CUDTagsRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchTagsResult struct {
	Create []product.Tag `json:"create"`
	Update []product.Tag `json:"update"`
	Delete []product.Tag `json:"delete"`
}

func (s productTagService) Batch(req CUDTagsRequest) (res BatchTagsResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("products/tags/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
