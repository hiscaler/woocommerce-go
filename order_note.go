package woocommerce

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type orderNoteService service

type OrderNotesQueryParams struct {
	queryParams
	Type string `url:"type,omitempty"`
}

func (m OrderNotesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Type, validation.When(m.Type != "", validation.In("any", "customer", "internal").Error("无效的类型"))),
	)
}

func (s orderNoteService) All(orderId int, params OrderNotesQueryParams) (items []entity.OrderNote, total, totalPages int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	urlValues, _ := query.Values(params)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get(fmt.Sprintf("/orders/%d/notes", orderId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
		total, totalPages, isLastPage = parseResponseTotal(params.Page, resp)
	}
	return
}

func (s orderNoteService) One(orderId, noteId int) (item entity.OrderNote, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/orders/%d/notes/%d", orderId, noteId))
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

func (s orderNoteService) Create(orderId int, req CreateOrderNoteRequest) (item entity.OrderNote, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
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

func (s orderNoteService) Delete(orderId, noteId int, force bool) (item entity.OrderNote, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/orders/%d/notes/%d", orderId, noteId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
