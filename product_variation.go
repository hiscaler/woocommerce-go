package woocommerce

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productVariationService service

// Product variations

type ProductVariationsQueryParams struct {
	queryParams
	Search string `json:"search,omitempty"`
}

func (m ProductVariationsQueryParams) Validate() error {
	return nil
}

// All List all product variations
func (s productVariationService) All(productId int, params ProductVariationsQueryParams) (items []entity.ProductVariation, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	urlValues, _ := query.Values(params)
	var res []entity.ProductVariation
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get(fmt.Sprintf("/products/%d/variations", productId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			items = res
		}
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}
