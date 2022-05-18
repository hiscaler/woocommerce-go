package entity

// SystemStatusSetting System status setting properties
type SystemStatusSetting struct {
	APIEnabled         bool     `json:"api_enabled"`
	ForceSSL           bool     `json:"force_ssl"`
	Currency           string   `json:"currency"`
	CurrencySymbol     string   `json:"currency_symbol"`
	CurrencyPosition   string   `json:"currency_position"`
	ThousandSeparator  string   `json:"thousand_separator"`
	DecimalSeparator   string   `json:"decimal_separator"`
	NumberOfDecimals   int      `json:"number_of_decimals"`
	GeolocationEnabled bool     `json:"geolocation_enabled"`
	Taxonomies         []string `json:"taxonomies"`
}
