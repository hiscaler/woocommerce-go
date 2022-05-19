package entity

type Taxes struct {
	ID               int        `json:"id"`
	RateCode         string     `json:"rate_code"`
	RateId           string     `json:"rate_id"`
	Label            string     `json:"label"`
	Compound         bool       `json:"compound"`
	TaxTotal         string     `json:"tax_total"`
	ShippingTaxTotal string     `json:"shipping_tax_total"`
	MetaData         []MetaData `json:"meta_data"`
}

type TaxLine struct {
	ID               int        `json:"id"`
	RateCode         string     `json:"rate_code"`
	RateId           string     `json:"rate_id"`
	Label            string     `json:"label"`
	Compound         bool       `json:"compound"`
	TaxTotal         string     `json:"tax_total"`
	ShippingTaxTotal string     `json:"shipping_tax_total"`
	MetaData         []MetaData `json:"meta_data"`
}
