package product

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity/product"
	jsoniter "github.com/json-iterator/go"
)

type TagsQueryParams struct {
	Context   string   `json:"context,omitempty"`
	Search    string   `json:"search,omitempty"`
	Exclude   []string `json:"exclude,omitempty"`
	Include   []string `json:"include,omitempty"`
	Offset    int      `json:"offset,omitempty"`
	Order     string   `json:"order,omitempty"`
	OrderBy   string   `json:"orderby,omitempty"`
	HideEmpty bool     `json:"hide_empty,omitempty"`
	Product   int      `json:"product,omitempty"`
	Slug      string   `json:"slug,omitempty"`
}

func (m TagsQueryParams) Validate() error {
	return nil
}

func (s service) Tags(params TagsQueryParams) (items []product.Tag, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []product.Tag
	qp := make(map[string]string, 0)
	resp, err := s.woo.Client.R().
		SetQueryParams(qp).
		Get("/products/tags")
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

func (s service) Tag(id int) (item product.Tag, err error) {
	var res product.Tag
	resp, err := s.woo.Client.R().Get(fmt.Sprintf("/products/tags/%d", id))
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

func (m UpsertProductTagRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name,
			validation.Required.Error("标签名称不能为空"),
		),
	)
}

func (s service) CreateTag(req UpsertProductTagRequest) (tag product.Tag, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.woo.Client.R().SetBody(req).Post("/products/tags")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &tag)
	}
	return
}

func (s service) UpdateTag(id int, req UpsertProductTagRequest) (tag product.Tag, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.woo.Client.R().SetBody(req).Put(fmt.Sprintf("/products/tags/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &tag)
	}
	return
}

func (s service) DeleteTag(id int) (tag product.Tag, err error) {
	resp, err := s.woo.Client.R().Delete(fmt.Sprintf("/products/tags/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &tag)
	}
	return
}
