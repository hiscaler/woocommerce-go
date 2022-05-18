package ordernote

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	OrderNotes(orderId int, params OrderNotesQueryParams) (items []entity.OrderNote, isLastPage bool, err error) // List all order notes
	OrderNote(orderId, noteId int) (item entity.OrderNote, err error)                                            // Retrieve an order note
	CreateOrderNote(orderId int, req CreateOrderNoteRequest) (item entity.OrderNote, err error)                  // Create an order note
	DeleteOrderNote(orderId, noteId int, force bool) (item entity.OrderNote, err error)                          // Delete an order note
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
