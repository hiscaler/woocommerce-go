package entity

// PaymentGateway payment gateway properties
type PaymentGateway struct {
	ID                int                              `json:"id"`
	Title             string                           `json:"title"`
	Description       string                           `json:"description"`
	Order             int                              `json:"order"`
	Enabled           bool                             `json:"enabled"`
	MethodTitle       string                           `json:"method_title"`
	MethodDescription string                           `json:"method_description"`
	MethodSupports    []string                         `json:"method_supports"`
	Settings          map[string]PaymentGatewaySetting `json:"settings"`
}

type PaymentGatewaySetting struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Default     string `json:"default"`
	Tip         string `json:"tip"`
	Placeholder string `json:"placeholder"`
}
