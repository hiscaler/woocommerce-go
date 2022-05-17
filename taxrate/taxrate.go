package taxrate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go"
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
	woocommerce.Query
	OrderBy string `url:"orderby,omitempty"`
	Class   string `url:"class,omitempty"`
}

func (m TaxRatesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "order", "priority").Error("无效的排序字段"))),
	)
}

func (s service) TaxRates(params TaxRatesQueryParams) (items []TaxRate, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []TaxRate
	params.TidyVars()
	urlValues, _ := query.Values(params)
	resp, err := s.woo.Client.R().SetQueryParamsFromValues(urlValues).Get("/taxes")
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
