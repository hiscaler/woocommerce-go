package woocommerce

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type orderRefundService service

// List all order refunds

type OrderRefundsQueryParams struct {
	queryParams
	Search        string `url:"search,omitempty"`
	After         string `url:"after,omitempty"`
	Before        string `url:"before,omitempty"`
	Exclude       []int  `url:"exclude,omitempty"`
	Include       []int  `url:"include,omitempty"`
	Parent        []int  `url:"parent,omitempty"`
	ParentExclude []int  `url:"parent_exclude,omitempty"`
	DecimalPoint  int    `url:"dp,omitempty"`
}

func (m OrderRefundsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Before, validation.When(m.Before != "", validation.By(func(value interface{}) error {
			dateStr, _ := value.(string)
			return IsValidateTime(dateStr)
		}))),
		validation.Field(&m.After, validation.When(m.After != "", validation.By(func(value interface{}) error {
			dateStr, _ := value.(string)
			return IsValidateTime(dateStr)
		}))),
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "date", "include", "title", "slug").Error("无效的排序值"))),
	)
}

func (s orderRefundService) All(orderId int, params OrderRefundsQueryParams) (items []entity.OrderRefund, total, totalPages int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	params.After = ToISOTimeString(params.After, false, true)
	params.Before = ToISOTimeString(params.Before, true, false)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(toValues(params)).Get(fmt.Sprintf("/orders/%d/refunds", orderId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
		total, totalPages, isLastPage = parseResponseTotal(params.Page, resp)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Create an order refund

type CreateOrderRefundRequest struct {
	Amount     float64                      `json:"amount,string"`
	Reason     string                       `json:"reason,omitempty"`
	RefundedBy int                          `json:"refunded_by,omitempty"`
	MetaData   []entity.Meta                `json:"meta_data,omitempty"`
	LineItems  []entity.OrderRefundLineItem `json:"line_items,omitempty"`
}

func (m CreateOrderRefundRequest) Validate() error {
	return nil
}

func (s orderRefundService) Create(orderId int, req CreateOrderRefundRequest) (item entity.OrderRefund, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
		SetBody(req).
		Post(fmt.Sprintf("/orders/%d/refunds", orderId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// One retrieve an order refund
func (s orderRefundService) One(orderId, refundId, decimalPoint int) (item entity.OrderRefund, err error) {
	if decimalPoint <= 0 || decimalPoint >= 6 {
		decimalPoint = 2
	}
	resp, err := s.httpClient.R().
		SetBody(map[string]int{"dp": decimalPoint}).
		Get(fmt.Sprintf("/orders/%d/refunds/%d", orderId, refundId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Delete delete an order refund
func (s orderRefundService) Delete(orderId, refundId int, force bool) (item entity.OrderRefund, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/orders/%d/refunds/%d", orderId, refundId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}
