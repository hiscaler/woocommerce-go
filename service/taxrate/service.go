package taxrate

import (
	"github.com/hiscaler/woocommerce-go"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	TaxRates(params TaxRatesQueryParams) (items []TaxRate, isLastPage bool, err error) // List all tax rates
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
