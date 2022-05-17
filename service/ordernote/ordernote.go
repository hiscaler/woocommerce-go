package ordernote

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/order"
	jsoniter "github.com/json-iterator/go"
)

type OrderNotesQueryParams struct {
	woocommerce.Query
	Type string `url:"type"`
}

func (m OrderNotesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Type, validation.When(m.Type != "", validation.In("any", "customer", "internal").Error("无效的类型"))),
	)
}

func (s service) OrderNotes(orderId int, params OrderNotesQueryParams) (items []order.Note, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	urlValues, _ := query.Values(params)
	var res []order.Note
	resp, err := s.woo.Client.R().SetQueryParamsFromValues(urlValues).Get(fmt.Sprintf("/orders/%d/notes", orderId))
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

func (s service) OrderNote(orderId, noteId int) (item order.Note, err error) {
	resp, err := s.woo.Client.R().Get(fmt.Sprintf("/orders/%d/notes/%d", orderId, noteId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
