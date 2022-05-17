package taxclass

import (
	"github.com/hiscaler/woocommerce-go"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	TaxClasses() (items []TaxClass, err error)                           // List all tax classes
	CreateTaxClass(req CreateTaxClassRequest) (item TaxClass, err error) // Create a tax class
	DeleteTaxClass(slug string) (item TaxClass, err error)               // Delete a tax class
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
