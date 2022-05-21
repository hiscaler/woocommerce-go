package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

// https://woocommerce.github.io/woocommerce-rest-api-docs/?php#coupon-properties

type couponService service

type CouponsQueryParams struct {
	queryParams
	Search  string `url:"search,omitempty"`
	After   string `url:"after,omitempty"`
	Before  string `url:"before,omitempty"`
	Exclude []int  `url:"exclude,omitempty"`
	Include []int  `url:"include,omitempty"`
	Code    string `url:"code,omitempty"`
}

func (m CouponsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "date", "title", "slug").Error("无效的排序字段"))),
	)
}

// All List all coupons
func (s couponService) All(params CouponsQueryParams) (items []entity.Coupon, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	resp, err := s.httpClient.R().SetQueryParamsFromValues(toValues(params)).Get("/coupons")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &items); err == nil {
			isLastPage = len(items) < params.PerPage
		}
	}
	return
}

// One Retrieve a coupon
func (s couponService) One(id int) (item entity.Coupon, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/coupons/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateCouponRequest struct {
	Code             string  `json:"code"`
	DiscountType     string  `json:"discount_type"`
	Amount           float64 `json:"amount,string"`
	IndividualUse    bool    `json:"individual_use"`
	ExcludeSaleItems bool    `json:"exclude_sale_items"`
	MinimumAmount    float64 `json:"minimum_amount,string"`
}

func (m CreateCouponRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.DiscountType, validation.In("percent", "fixed_cart", "fixed_product").Error("无效的折扣类型")),
		validation.Field(&m.Amount,
			validation.Min(0.0).Error("金额不能小于 {{.threshold}}"),
			validation.When(m.DiscountType == "percent", validation.Max(100.0).Error("折扣比例不能大于 {{.threshold}}")),
		),
		validation.Field(&m.MinimumAmount, validation.Min(0.0).Error("最小金额不能小于 {{.threshold}}")),
	)
}

// Create Create a coupon
func (s couponService) Create(req CreateCouponRequest) (item entity.Coupon, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/coupons")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

type UpdateCouponRequest struct {
	Code             string  `json:"code,omitempty"`
	DiscountType     string  `json:"discount_type,omitempty"`
	Amount           float64 `json:"amount,omitempty,string"`
	IndividualUse    bool    `json:"individual_use,omitempty"`
	ExcludeSaleItems bool    `json:"exclude_sale_items,omitempty"`
	MinimumAmount    float64 `json:"minimum_amount,omitempty,string"`
}

func (m UpdateCouponRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.DiscountType, validation.When(m.DiscountType != "", validation.In("percent", "fixed_cart", "fixed_product").Error("无效的折扣类型"))),
		validation.Field(&m.Amount,
			validation.Min(0.0).Error("金额不能小于 {{.threshold}}"),
			validation.When(m.DiscountType == "percent", validation.Max(100.0).Error("折扣比例不能大于 {{.threshold}}")),
		),
		validation.Field(&m.MinimumAmount, validation.Min(0.0).Error("最小金额不能小于 {{.threshold}}")),
	)
}

// Update Update a coupon
func (s couponService) Update(id int, req UpdateCouponRequest) (item entity.Coupon, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/coupons/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete a coupon

func (s couponService) Delete(id int, force bool) (item entity.Coupon, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/coupons/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch update coupons

type BatchCouponsCreateItem = CreateCouponRequest
type BatchCouponsUpdateItem struct {
	ID string `json:"id"`
	BatchCouponsCreateItem
}

type BatchCouponsRequest struct {
	Create []BatchCouponsCreateItem `json:"create,omitempty"`
	Update []BatchCouponsUpdateItem `json:"update,omitempty"`
	Delete []int                    `json:"delete,omitempty"`
}

func (m BatchCouponsRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchCouponsResult struct {
	Create []entity.Coupon `json:"create"`
	Update []entity.Coupon `json:"update"`
	Delete []entity.Coupon `json:"delete"`
}

// Batch Batch create/update/delete coupons
func (s couponService) Batch(req BatchCouponsRequest) (res BatchCouponsResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/coupons/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
