package order

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/order"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Orders(params OrdersQueryParams) (items []order.Order, isLastPage bool, err error) // List all orders
	Order(id int) (item order.Order, err error)                                        // Retrieve an order
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
