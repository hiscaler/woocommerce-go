package order

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Orders(params OrdersQueryParams) (items []entity.Order, isLastPage bool, err error) // List all orders
	Order(id int) (item entity.Order, err error)                                        // Retrieve an order
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
