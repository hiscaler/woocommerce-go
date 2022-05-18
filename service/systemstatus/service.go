package systemstatus

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	SystemStatus() (item entity.SystemStatus, err error) // List all system status items
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
