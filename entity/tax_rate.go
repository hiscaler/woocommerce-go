package entity

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
