package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

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
		validation.Field(&m.Email, validation.When(m.Email != "", is.EmailFormat.Error("无效的邮箱"))),
		validation.Field(&m.FirstName, validation.Required.Error("姓不能为空")),
		validation.Field(&m.LastName, validation.Required.Error("名不能为空")),
	)
}
