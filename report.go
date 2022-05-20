package woocommerce

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/constant"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type reportService service

type ReportsQueryParams struct {
	Context string `url:"context,omitempty"`
	Period  string `url:"period,omitempty"`
	DateMin string `url:"date_min,omitempty"`
	DateMax string `url:"date_max,omitempty"`
}

func (m ReportsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Period, validation.When(m.Period != "", validation.In("week", "month", "last_month", "year").Error("无效的报表周期"))),
		validation.Field(&m.DateMin,
			validation.Required.Error("报表开始时间不能为空"),
			validation.Date(constant.DateFormat).Error("报表开始时间格式无效"),
		),
		validation.Field(&m.DateMax,
			validation.Required.Error("报表结束时间不能为空"),
			validation.Date(constant.DateFormat).Error("报表结束时间格式无效"),
			validation.By(func(value interface{}) (err error) {
				dateMin, err := time.Parse(constant.DateFormat, m.DateMin)
				if err != nil {
					return
				}
				dateMax, err := time.Parse(constant.DateFormat, m.DateMax)
				if err != nil {
					return
				}
				if dateMax.Before(dateMin) {
					return fmt.Errorf("结束时间不能小于 %s", m.DateMin)
				}
				return nil
			}),
		),
	)
}

// All list all reports
func (s reportService) All() (items []entity.Report, err error) {
	resp, err := s.httpClient.R().Get("/reports")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// Sales reports

type SalesReportsQueryParams = ReportsQueryParams

// SalesReports list all sales reports
func (s reportService) SalesReports(req SalesReportsQueryParams) (items []entity.SaleReport, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().Get("/reports/sales")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// TopSellerReports list all sales reports

type TopSellerReportsQueryParams = SalesReportsQueryParams

func (s reportService) TopSellerReports(req SalesReportsQueryParams) (items []entity.TopSellerReport, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().Get("/reports/top_sellers")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// CouponTotals retrieve coupons totals
func (s reportService) CouponTotals() (items []entity.CouponTotal, err error) {
	resp, err := s.httpClient.R().Get("/reports/coupons/totals")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// CustomerTotals retrieve customer totals
func (s reportService) CustomerTotals() (items []entity.CustomerTotal, err error) {
	resp, err := s.httpClient.R().Get("/reports/customers/totals")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// OrderTotals retrieve customer totals
func (s reportService) OrderTotals() (items []entity.OrderTotal, err error) {
	resp, err := s.httpClient.R().Get("/reports/orders/totals")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// ProductTotals retrieve product totals
func (s reportService) ProductTotals() (items []entity.OrderTotal, err error) {
	resp, err := s.httpClient.R().Get("/reports/products/totals")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// ReviewTotals retrieve review totals
func (s reportService) ReviewTotals() (items []entity.OrderTotal, err error) {
	resp, err := s.httpClient.R().Get("/reports/reviews/totals")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}
