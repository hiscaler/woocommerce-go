package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productTagService service

type ProductTagsQueryParams struct {
	queryParams
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

func (s productTagService) All(params ProductTagsQueryParams) (items []entity.ProductTag, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []entity.ProductTag
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

func (s productTagService) One(id int) (item entity.ProductTag, err error) {
	var res entity.ProductTag
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

type UpsertProductTagRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateProductTagRequest = UpsertProductTagRequest
type UpdateProductTagRequest = UpsertProductTagRequest

func (m UpsertProductTagRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name,
			validation.Required.Error("标签名称不能为空"),
		),
	)
}

func (s productTagService) Create(req CreateProductTagRequest) (item entity.ProductTag, err error) {
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

func (s productTagService) Update(id int, req UpdateProductTagRequest) (item entity.ProductTag, err error) {
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

func (s productTagService) Delete(id int, force bool) (item entity.ProductTag, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/tags/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch tag create,update and delete operation

type BatchProductTagsCreateItem = UpsertProductTagRequest

type BatchProductTagsUpdateItem struct {
	ID int `json:"id"`
	UpsertProductTagRequest
}
type BatchProductTagsRequest struct {
	Create []BatchProductTagsCreateItem `json:"create,omitempty"`
	Update []BatchProductTagsUpdateItem `json:"update,omitempty"`
	Delete []int                        `json:"delete,omitempty"`
}

func (m BatchProductTagsRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchProductTagsResult struct {
	Create []entity.ProductTag `json:"create"`
	Update []entity.ProductTag `json:"update"`
	Delete []entity.ProductTag `json:"delete"`
}

func (s productTagService) Batch(req BatchProductTagsRequest) (res BatchProductTagsResult, err error) {
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
