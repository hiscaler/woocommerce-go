package entity

// Customer customer properties
type Customer struct {
	ID               int      `json:"id"`
	DateCreated      string   `json:"date_created"`
	DateCreatedGMT   string   `json:"date_created_gmt"`
	DateModified     string   `json:"date_modified"`
	DateModifiedGMT  string   `json:"date_modified_gmt"`
	Email            string   `json:"email"`
	FirstName        string   `json:"first_name"`
	LastName         string   `json:"last_name"`
	Role             string   `json:"role"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	Billing          Billing  `json:"billing"`
	Shipping         Shipping `json:"shipping"`
	IsPayingCustomer bool     `json:"is_paying_customer"`
	AvatarURL        string   `json:"avatar_url"`
	MetaData         []Meta   `json:"meta_data"`
}
