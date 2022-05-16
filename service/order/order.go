package order

import (
	"fmt"
	"github.com/hiscaler/woocommerce-go/entity/order"
	jsoniter "github.com/json-iterator/go"
)

type OrdersQueryParams struct {
	Context       string   `json:"context,omitempty"`
	Search        string   `json:"search"`
	After         string   `json:"after"`
	Before        string   `json:"before"`
	Exclude       []int    `json:"exclude"`
	Include       []int    `json:"include"`
	Offset        int      `json:"offset"`
	Order         string   `json:"order,omitempty"`
	OrderBy       string   `json:"Orderby,omitempty"`
	Parent        []int    `json:"parent"`
	ParentExclude []int    `json:"parent_exclude"`
	Status        []string `json:"status,omitempty"`
	Customer      int      `json:"customer"`
	Product       int      `json:"product"`
	DecimalPoint  int      `json:"dp,omitempty"`
}

func (m OrdersQueryParams) Validate() error {
	return nil
}

func (s service) Orders(params OrdersQueryParams) (items []order.Order, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []order.Order
	qp := make(map[string]string, 0)
	resp, err := s.woo.Client.R().
		SetQueryParams(qp).
		Get("/orders")
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
