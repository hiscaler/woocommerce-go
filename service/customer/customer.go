package customer

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/customer"
	jsoniter "github.com/json-iterator/go"
)

type CustomersQueryParams struct {
	woocommerce.Query
	Search  string `url:"search"`
	Exclude []int  `url:"exclude"`
	Include []int  `url:"include"`
	Email   string `url:"email"`
	Role    string `url:"role"`
}

func (m CustomersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "name", "registered_date").Error("无效的排序字段"))),
		validation.Field(&m.Role, validation.When(m.Role != "", validation.In("all", "administrator", "editor", "author", "contributor", "subscriber", "shop_manager").Error("无效的角色"))),
	)
}

func (s service) Customers(params CustomersQueryParams) (items []customer.Customer, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	urlValues, _ := query.Values(params)
	var res []customer.Customer
	resp, err := s.woo.Client.R().SetQueryParamsFromValues(urlValues).Get("/customers")
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

func (s service) Customer(id int) (item customer.Customer, err error) {
	resp, err := s.woo.Client.R().Get(fmt.Sprintf("/customers/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
