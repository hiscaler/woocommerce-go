package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type orderService service

type OrdersQueryParams struct {
	queryParams
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

// All list all orders
func (s orderService) All(params OrdersQueryParams) (items []entity.Order, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	resp, err := s.httpClient.R().SetQueryParamsFromValues(toValues(params)).Get("/orders")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &items); err == nil {
			isLastPage = len(items) < params.PerPage
		}
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// One retrieve an order
func (s orderService) One(id int) (item entity.Order, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/orders/%d", id))
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

// Create order

type CreateOrderRequest struct {
	Status             string                `json:"status,omitempty"`
	Currency           string                `json:"currency,omitempty"`
	CurrencySymbol     string                `json:"currency_symbol,omitempty"`
	PricesIncludeTax   bool                  `json:"prices_include_tax,omitempty"`
	CustomerId         int                   `json:"customer_id,omitempty"`
	CustomerNote       string                `json:"customer_note,omitempty"`
	Billing            *entity.Billing       `json:"billing,omitempty"`
	Shipping           *entity.Shipping      `json:"shipping,omitempty"`
	PaymentMethod      string                `json:"payment_method,omitempty"`
	PaymentMethodTitle string                `json:"payment_method_title,omitempty"`
	TransactionId      string                `json:"transaction_id,omitempty"`
	MetaData           []entity.MetaData     `json:"meta_data,omitempty"`
	LineItems          []entity.LineItem     `json:"line_items,omitempty"`
	TaxLines           []entity.TaxLine      `json:"tax_lines,omitempty"`
	ShippingLines      []entity.ShippingLine `json:"shipping_lines,omitempty"`
	FeeLines           []entity.FeeLine      `json:"fee_lines,omitempty"`
	CouponLines        []entity.CouponLine   `json:"coupon_lines,omitempty"`
	SetPaid            bool                  `json:"set_paid,omitempty"`
}

func (m CreateOrderRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Status, validation.When(m.Status != "", validation.In("pending", "processing", "on-hold", "completed", "cancelled", "refunded", "failed", "trash").Error("无效的状态"))),
	)
}

func (s orderService) Create(req CreateOrderRequest) (item entity.Order, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/orders")
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

// Update order

type UpdateOrderRequest = CreateOrderRequest

func (s orderService) Update(id int, req UpdateOrderRequest) (item entity.Order, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/orders/%d", id))
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

// Delete delete an order
func (s orderService) Delete(id int, force bool) (item entity.Order, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/orders/%d", id))
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
