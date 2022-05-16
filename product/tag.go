package product

import (
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
