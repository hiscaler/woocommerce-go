package customer

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/customer"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Customers(params CustomersQueryParams) (items []customer.Customer, isLastPage bool, err error) // List all customers
	Customer(id int) (item customer.Customer, err error)                                           // Retrieve a customer
	CreateCustomer(req CreateCustomerRequest) (item customer.Customer, err error)                                     // Create a customer
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
