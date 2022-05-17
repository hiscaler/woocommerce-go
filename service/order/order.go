package order

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/order"
	jsoniter "github.com/json-iterator/go"
)

type OrdersQueryParams struct {
	woocommerce.Query
	Search        string   `url:"search"`
	After         string   `url:"after"`
	Before        string   `url:"before"`
	Exclude       []int    `url:"exclude"`
	Include       []int    `url:"include"`
	Parent        []int    `url:"parent"`
	ParentExclude []int    `url:"parent_exclude"`
	Status        []string `url:"status,omitempty"`
	Customer      int      `url:"customer"`
	Product       int      `url:"product"`
	DecimalPoint  int      `url:"dp,omitempty"`
}

func (m OrdersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "date", "include", "title", "slug").Error("无效的排序字段"))),
		validation.Field(&m.Status, validation.When(m.OrderBy != "", validation.In("any", "pending", "processing", "on-hold", "completed", "cancelled", "refunded", "failed ", "trash").Error("无效的状态"))),
	)
}

func (s service) Orders(params OrdersQueryParams) (items []order.Order, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	urlValues, _ := query.Values(params)
	var res []order.Order
	resp, err := s.woo.Client.R().SetQueryParamsFromValues(urlValues).Get("/orders")
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

func (s service) Order(id int) (item order.Order, err error) {
	var res order.Order
	resp, err := s.woo.Client.R().Get(fmt.Sprintf("/orders/%d", id))
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
