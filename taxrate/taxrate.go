package taxrate

import jsoniter "github.com/json-iterator/go"

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
}

func (m TaxRatesQueryParams) Validate() error {
	return nil
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
