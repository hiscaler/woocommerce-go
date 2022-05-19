package entity

// ShippingZoneMethod shipping zone method properties
type ShippingZoneMethod = ShippingMethod

// ShippingZoneMethodSetting shipping zone method setting properties
type ShippingZoneMethodSetting struct {
	ID          int    `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Default     string `json:"default"`
	Tip         string `json:"tip"`
	PlaceHolder string `json:"place_holder"`
}
