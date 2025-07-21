package woocommerce

import (
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type webhookService service

type WebhooksQueryParams struct {
	queryParams
	Search  string `url:"search"`
	After   string `url:"after"`
	Before  string `url:"before"`
	Exclude []int  `url:"exclude"`
	Include []int  `url:"include"`
	Status  string `url:"status"`
}

func (m WebhooksQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Before, validation.When(m.Before != "", validation.By(func(value interface{}) error {
			dateStr, _ := value.(string)
			return IsValidateTime(dateStr)
		}))),
		validation.Field(&m.After, validation.When(m.After != "", validation.By(func(value interface{}) error {
			dateStr, _ := value.(string)
			return IsValidateTime(dateStr)
		}))),
	)
}

// All List all webhooks
func (s webhookService) All(params WebhooksQueryParams) (items []entity.Webhook, total, totalPages int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	params.After = ToISOTimeString(params.After, false, true)
	params.Before = ToISOTimeString(params.Before, true, false)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(toValues(params)).Get("/products/webhooks")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
		total, totalPages, isLastPage = parseResponseTotal(params.Page, resp)
	}
	return
}

// One Retrieve a webhook
func (s webhookService) One(id int) (item entity.Webhook, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/webhooks/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateWebhookRequest struct {
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	Topic       string `json:"topic,omitempty"`
	DeliveryURL string `json:"delivery_url,omitempty"`
	Secret      string `json:"secret,omitempty"`
}

func (m CreateWebhookRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.DeliveryURL, validation.When(m.DeliveryURL != "", is.URL.Error("投递 URL 格式错误"))),
		validation.Field(&m.Status, validation.When(m.Status != "", validation.In("active", "paused", "disabled").Error("无效的状态值"))),
	)
}

// Create Create a product attribute
func (s webhookService) Create(req CreateWebhookRequest) (item entity.Webhook, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/webhooks")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

type UpdateWebhookRequest = CreateWebhookRequest

// Update Update a webhook
func (s webhookService) Update(id int, req UpdateWebhookRequest) (item entity.Webhook, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/products/webhooks/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete a webhook

func (s webhookService) Delete(id int, force bool) (item entity.Webhook, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/webhooks/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch update webhooks

type BatchWebhooksCreateItem = CreateWebhookRequest
type BatchWebhooksUpdateItem struct {
	ID string `json:"id"`
	BatchWebhooksCreateItem
}

type BatchWebhooksRequest struct {
	Create []BatchWebhooksCreateItem `json:"create,omitempty"`
	Update []BatchWebhooksUpdateItem `json:"update,omitempty"`
	Delete []int                     `json:"delete,omitempty"`
}

func (m BatchWebhooksRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchWebhooksResult struct {
	Create []entity.Webhook `json:"create"`
	Update []entity.Webhook `json:"update"`
	Delete []entity.Webhook `json:"delete"`
}

// Batch Batch create/update/delete webhooks
func (s webhookService) Batch(req BatchWebhooksRequest) (res BatchWebhooksResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/webhooks/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
