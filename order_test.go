package woocommerce

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Query orders
func ExampleAll() {
	params := OrdersQueryParams{
		After: "2022-06-10",
	}
	params.PerPage = 100
	for {
		orders, total, totalPages, isLastPage, err := wooClient.Services.Order.All(params)
		if err != nil {
			break
		}
		fmt.Println(fmt.Sprintf("Page %d/%d", total, totalPages))
		// read orders
		for _, order := range orders {
			_ = order
		}
		if err != nil || isLastPage {
			break
		}
		params.Page++
	}
}

func TestOrderService_All(t *testing.T) {
	params := OrdersQueryParams{
		After: "2022-06-10",
	}
	params.PerPage = 100
	items, _, _, isLastPage, err := wooClient.Services.Order.All(params)
	if err != nil {
		t.Fatalf("wooClient.Services.Order.All error: %s", err.Error())
	}
	if len(items) > 0 {
		orderId = items[0].ID
	}
	assert.Equal(t, true, isLastPage, "check isLastPage")
}

func TestOrderService_AllByArrayParams(t *testing.T) {
	params := OrdersQueryParams{
		Status:  []string{"completed"},
		Include: []int{914, 849},
	}
	params.PerPage = 300
	_, _, _, isLastPage, err := wooClient.Services.Order.All(params)
	if err != nil {
		t.Fatalf("wooClient.Services.Order.All By Array Params error: %s", err.Error())
	}
	assert.Equal(t, true, isLastPage, "check isLastPage")
}

func TestOrderService_One(t *testing.T) {
	item, err := wooClient.Services.Order.One(orderId)
	if err != nil {
		t.Errorf("wooClient.Services.Order.One(%d) error: %s", orderId, err.Error())
	} else {
		assert.Equal(t, orderId, item.ID, "order id")
	}
}

func TestOrderService_Create(t *testing.T) {
	req := CreateOrderRequest{}
	item, err := wooClient.Services.Order.Create(req)
	if err != nil {
		t.Fatalf("wooClient.Services.Order.Create error: %s", err.Error())
	}
	orderId = item.ID
}

func TestOrderService_Update(t *testing.T) {
	t.Run("getOrderId", getOrderId)
	req := UpdateOrderRequest{
		PaymentMethod:      "paypal",
		PaymentMethodTitle: "Paypal",
	}
	item, err := wooClient.Services.Order.Update(orderId, req)
	if err != nil {
		t.Fatalf("wooClient.Services.Order.Update error: %s", err.Error())
	} else {
		assert.Equal(t, orderId, item.ID, "order id")
	}
}

func TestOrderService_Delete(t *testing.T) {
	_, err := wooClient.Services.Order.Delete(orderId, true)
	if err != nil {
		t.Fatalf("wooClient.Services.Order.Delete(%d, true) error: %s", orderId, err.Error())
	}
}
