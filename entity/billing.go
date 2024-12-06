package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Billing order billing properties
type Billing struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Address1  string `json:"address_1,omitempty"`
	Address2  string `json:"address_2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Postcode  string `json:"postcode,omitempty"`
	Country   string `json:"country,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func (m Billing) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.When(m.Email != "", is.EmailFormat.Error("invalid email"))),
		validation.Field(&m.FirstName, validation.Required.Error("first name cannot be empty")),
		validation.Field(&m.LastName, validation.Required.Error("last name cannot be empty")),
	)
}
