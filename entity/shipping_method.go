package entity

type ShippingMethod struct {
	InstanceId        int                         `json:"instance_id"`
	Title             string                      `json:"title"`
	Order             int                         `json:"order"`
	Enabled           bool                        `json:"enabled"`
	MethodId          string                      `json:"method_id"`
	MethodTitle       string                      `json:"method_title"`
	MethodDescription string                      `json:"method_description"`
	Settings          []ShippingZoneMethodSetting `json:"settings"`
}
