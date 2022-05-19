package entity

// ShippingZone shipping zone properties
type ShippingZone struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}
