package ordernote

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/order"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	OrderNotes(orderId int, params OrderNotesQueryParams) (items []order.Note, isLastPage bool, err error) // List all order notes
	OrderNote(orderId, noteId int) (item order.Note, err error)                                            // Retrieve an order note
	CreateOrderNote(orderId int, req CreateOrderNoteRequest) (item order.Note, err error)                  // Create an order note
	DeleteOrderNote(orderId, noteId int, force bool) (item order.Note, err error)                          // Delete an order note
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
