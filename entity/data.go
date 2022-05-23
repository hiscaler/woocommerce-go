package entity

// Data data properties
type Data struct {
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

// Continent continent properties
type Continent struct {
	Code      string             `json:"code"`
	Name      string             `json:"name"`
	Countries []ContinentCountry `json:"countries"` // Only code, name, []state?
}

// ContinentCountry continent country properties
type ContinentCountry struct {
	Code          string  `json:"code"`
	CurrencyCode  string  `json:"currency_code"`
	CurrencyPos   string  `json:"currency_pos"`
	DecimalSep    string  `json:"decimal_sep"`
	DimensionUnit string  `json:"dimension_unit"`
	Name          string  `json:"name"`
	NumDecimals   int     `json:"num_decimals"`
	States        []State `json:"states"`
	ThousandSep   string  `json:"thousand_sep"`
	WeightUnit    string  `json:"weight_unit"`
}

// State state properties
type State struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Country country properties
type Country struct {
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	States []State `json:"states"`
}

// Currency currency properties
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
