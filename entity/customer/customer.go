package customer

import (
	"github.com/hiscaler/woocommerce-go/entity"
)

type Customer struct {
	ID               int               `json:"id"`
	DateCreated      string            `json:"date_created"`
	DateCreatedGMT   string            `json:"date_created_gmt"`
	DateModified     string            `json:"date_modified"`
	DateModifiedGMT  string            `json:"date_modified_gmt"`
	Email            string            `json:"email"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	Role             string            `json:"role"`
	Username         string            `json:"username"`
	Password         string            `json:"password"`
	Billing          entity.Billing    `json:"billing"`
	Shipping         entity.Shipping   `json:"shipping"`
	IsPayingCustomer bool              `json:"is_paying_customer"`
	AvatarURL        string            `json:"avatar_url"`
	MetaData         []entity.MetaData `json:"meta_data"`
}
