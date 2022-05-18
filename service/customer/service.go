package customer

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Customers(params CustomersQueryParams) (items []entity.Customer, isLastPage bool, err error) // List all customers
	Customer(id int) (item entity.Customer, err error)                                           // Retrieve a customer
	CreateCustomer(req CreateCustomerRequest) (item entity.Customer, err error)                  // Create a customer
	UpdateCustomer(id int, req UpdateCustomerRequest) (item entity.Customer, err error)          // Update a customer
	DeleteCustomer(id int) (item entity.Customer, err error)                                     // Delete a customer
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
