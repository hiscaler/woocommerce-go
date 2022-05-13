package entity

type Shipping struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
}

type ShippingLine struct {
	Id          int        `json:"id"`
	MethodTitle string     `json:"method_title"`
	MethodId    string     `json:"method_id"`
	Total       string     `json:"total"`
	TotalTax    string     `json:"total_tax"`
	Taxes       []Taxes    `json:"taxes"`
	MetaData    []MetaData `json:"meta_data"`
}
