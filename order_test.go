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
		t.Errorf("wooClient.Services.Order.All error: %s", err.Error())
	} else {
		t.Logf("items = %#v", jsonx.ToPrettyJson(items))
	}
}

func TestOrderService_One(t *testing.T) {
	orderId := 849
	item, err := wooClient.Services.Order.One(orderId)
	if err != nil {
		t.Errorf("wooClient.Services.Order.One(%d) error: %s", orderId, err.Error())
	} else {
		assert.Equal(t, orderId, item.ID, "order id")
	}
}
