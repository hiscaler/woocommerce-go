package woocommerce

import (
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productReviewService service

type ProductReviewsQueryParams struct {
	queryParams
	Search          string   `url:"search,omitempty"`
	After           string   `url:"after,omitempty"`
	Before          string   `url:"before,omitempty"`
	Exclude         []int    `url:"exclude,omitempty"`
	Include         []int    `url:"include,omitempty"`
	Reviewer        []int    `url:"reviewer,omitempty"`
	ReviewerExclude []int    `url:"reviewer_exclude,omitempty"`
	ReviewerEmail   []string `url:"reviewer_email,omitempty"`
	Product         []int    `url:"product,omitempty"`
	Status          string   `url:"status,omitempty"`
}

func (m ProductReviewsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Before, validation.When(m.Before != "", validation.By(func(value interface{}) error {
			dateStr, _ := value.(string)
			return IsValidateTime(dateStr)
		}))),
		validation.Field(&m.After, validation.When(m.After != "", validation.By(func(value interface{}) error {
			dateStr, _ := value.(string)
			return IsValidateTime(dateStr)
		}))),
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "date", "date_gmt", "slug", "include", "product").Error("invalid sort method"))),
	)
}

// All List all product reviews
func (s productReviewService) All(params ProductReviewsQueryParams) (items []entity.ProductReview, total, totalPages int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	params.After = ToISOTimeString(params.After, false, true)
	params.Before = ToISOTimeString(params.Before, true, false)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(toValues(params)).Get("/products/reviews")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
		total, totalPages, isLastPage = parseResponseTotal(params.Page, resp)
	}
	return
}

// One Retrieve a product review
func (s productReviewService) One(id int) (item entity.ProductReview, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/reviews/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateProductReviewRequest struct {
	ProductId     int    `json:"product_id,omitempty"`
	Status        string `json:"status,omitempty"`
	Reviewer      string `json:"reviewer,omitempty"`
	ReviewerEmail string `json:"reviewer_email,omitempty"`
	Review        string `json:"review,omitempty"`
	Rating        int    `json:"rating,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
}

func (m CreateProductReviewRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Status, validation.When(m.Status != "", validation.In("approved", "hold", "spam", "unspam", "trash", "untrash").Error("invalid status"))),
		validation.Field(&m.Rating,
			validation.Min(0).Error("rating minimum is {{threshold}}"),
			validation.Min(5).Error("rating maximum is {{threshold}}"),
		),
	)
}

// Create Create a product review
func (s productReviewService) Create(req CreateProductReviewRequest) (item entity.ProductReview, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/reviews")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

type UpdateProductReviewRequest = CreateProductReviewRequest

// Update Update a product review
func (s productReviewService) Update(id int, req UpdateProductReviewRequest) (item entity.ProductReview, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/products/reviews/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete a product review

func (s productReviewService) Delete(id int, force bool) (item entity.ProductReview, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/reviews/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch update product reviews

type BatchProductReviewsCreateItem = CreateProductReviewRequest
type BatchProductReviewsUpdateItem struct {
	ID string `json:"id"`
	BatchProductReviewsCreateItem
}

type BatchProductReviewsRequest struct {
	Create []BatchProductReviewsCreateItem `json:"create,omitempty"`
	Update []BatchProductReviewsUpdateItem `json:"update,omitempty"`
	Delete []int                           `json:"delete,omitempty"`
}

func (m BatchProductReviewsRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("invalid request data")
	}
	return nil
}

type BatchProductReviewsResult struct {
	Create []entity.ProductReview `json:"create"`
	Update []entity.ProductReview `json:"update"`
	Delete []entity.ProductReview `json:"delete"`
}

// Batch Batch create/update/delete product reviews
func (s productReviewService) Batch(req BatchProductReviewsRequest) (res BatchProductReviewsResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/reviews/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
