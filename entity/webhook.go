package entity

// Webhook webhook properties
type Webhook struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Status          string   `json:"status"`
	Topic           string   `json:"topic"`
	Resource        string   `json:"resource"`
	Event           string   `json:"event"`
	Hooks           []string `json:"hooks"`
	DeliveryURL     string   `json:"delivery_url"`
	Secret          string   `json:"secret"`
	DateCreated     string   `json:"date_created"`
	DateCreatedGMT  string   `json:"date_created_gmt"`
	DateModified    string   `json:"date_modified"`
	DateModifiedGMT string   `json:"date_modified_gmt"`
}
