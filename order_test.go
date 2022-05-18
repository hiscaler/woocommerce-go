package woocommerce

import (
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderService_All(t *testing.T) {
	params := OrdersQueryParams{}
	items, _, err := wooClient.Services.Order.All(params)
	if err != nil {
		t.Fatalf("wooClient.Services.Order.All error: %s", err.Error())
	} else {
		t.Logf("items = %#v", jsonx.ToPrettyJson(items))
		if len(items) > 0 {
			orderId = items[0].ID
		}
	}
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

func TestOrderService_Delete(t *testing.T) {
	_, err := wooClient.Services.Order.Delete(orderId, true)
	if err != nil {
		t.Fatalf("wooClient.Services.Order.Delete(%d, true) error: %s", orderId, err.Error())
	}
}
