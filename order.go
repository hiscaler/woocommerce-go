package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type orderService service

type OrdersQueryParams struct {
	QueryParams
	Search        string   `url:"search,omitempty"`
	After         string   `url:"after,omitempty"`
	Before        string   `url:"before,omitempty"`
	Exclude       []int    `url:"exclude,omitempty"`
	Include       []int    `url:"include,omitempty"`
	Parent        []int    `url:"parent,omitempty"`
	ParentExclude []int    `url:"parent_exclude,omitempty"`
	Status        []string `url:"status,omitempty,omitempty"`
	Customer      int      `url:"customer,omitempty"`
	Product       int      `url:"product,omitempty"`
	DecimalPoint  int      `url:"dp,omitempty"`
}

func (m OrdersQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "date", "include", "title", "slug").Error("无效的排序字段"))),
		validation.Field(&m.Status, validation.When(len(m.Status) > 0, validation.By(func(value interface{}) error {
			statuses, ok := value.([]string)
			if !ok {
				return errors.New("无效的状态值")
			}
			validStatuses := []string{"any", "pending", "processing", "on-hold", "completed", "cancelled", "refunded", "failed ", "trash"}
			for _, status := range statuses {
				valid := false
				for _, validStatus := range validStatuses {
					if status == validStatus {
						valid = true
					}
				}
				if !valid {
					return fmt.Errorf("无效的状态值：%s", status)
				}
			}
			return nil
		}))),
	)
}

func (s orderService) All(params OrdersQueryParams) (items []entity.Order, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	urlValues, _ := query.Values(params)
	var res []entity.Order
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get("/orders")
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

func (s orderService) One(id int) (item entity.Order, err error) {
	var res entity.Order
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/orders/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			item = res
		}
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}
