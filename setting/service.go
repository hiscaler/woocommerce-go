package setting

import (
	"github.com/hiscaler/woocommerce-go"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Groups() (items []Group, err error) // List all settings groups
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
