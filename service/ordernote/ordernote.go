package ordernote

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

type OrderNotesQueryParams struct {
	woocommerce.QueryParams
	Type string `url:"type"`
}

func (m OrderNotesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Type, validation.When(m.Type != "", validation.In("any", "customer", "internal").Error("无效的类型"))),
	)
}

func (s service) OrderNotes(orderId int, params OrderNotesQueryParams) (items []entity.OrderNote, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	urlValues, _ := query.Values(params)
	var res []entity.OrderNote
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

func (s service) OrderNote(orderId, noteId int) (item entity.OrderNote, err error) {
	resp, err := s.woo.Client.R().Get(fmt.Sprintf("/orders/%d/notes/%d", orderId, noteId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create order note

type CreateOrderNoteRequest struct {
	Note string `json:"note"`
}

func (m CreateOrderNoteRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Note, validation.Required.Error("内容不能为空")),
	)
}

func (s service) CreateOrderNote(orderId int, req CreateOrderNoteRequest) (item entity.OrderNote, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.woo.Client.R().
		SetBody(req).
		Post(fmt.Sprintf("/orders/%d/notes", orderId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s service) DeleteOrderNote(orderId, noteId int, force bool) (item entity.OrderNote, err error) {
	resp, err := s.woo.Client.R().
		SetBody(map[string]string{
			"force": strconv.FormatBool(force),
		}).
		Delete(fmt.Sprintf("/orders/%d/notes/%d", orderId, noteId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
