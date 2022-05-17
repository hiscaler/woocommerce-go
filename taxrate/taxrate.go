package taxrate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	jsoniter "github.com/json-iterator/go"
)

// TaxRate tax rate properites
type TaxRate struct {
	ID        int      `json:"id"`
	Country   string   `json:"country"`
	State     string   `json:"state"`
	Postcode  string   `json:"postcode"`
	City      string   `json:"city"`
	Postcodes []string `json:"postcodes"`
	Cities    []string `json:"cities"`
	Rate      string   `json:"rate"`
	Name      string   `json:"name"`
	Priority  int      `json:"priority"`
	Compound  bool     `json:"compound"`
	Shipping  bool     `json:"shipping"`
	Order     int      `json:"order"`
	Class     string   `json:"class"`
}

type TaxRatesQueryParams struct {
	Context string `url:"context"`
	Order   string `url:"order,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	OrderBy string `url:"orderby,omitempty"`
	Class   string `url:"class,omitempty"`
}

func (m TaxRatesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Context, validation.In("view", "edit").Error("错误的请求范围")),
		validation.Field(&m.Order, validation.When(m.Order != "", validation.In("asc", "desc").Error("错误的排序方式"))),
	)
}

func (s service) TaxRates(params TaxRatesQueryParams) (items []TaxRate, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []TaxRate
	qp := make(map[string]string, 0)
	resp, err := s.woo.Client.R().
		SetQueryParams(qp).
		Get("/taxes")
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
