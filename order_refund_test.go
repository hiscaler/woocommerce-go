package woocommerce

import (
	"errors"
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderRefundService_All(t *testing.T) {
	t.Run("getOrderId", getOrderId)
	params := OrderRefundsQueryParams{}
	items, _, _, _, err := wooClient.Services.OrderRefund.All(orderId, params)
	if err != nil {
		t.Errorf("wooClient.Services.OrderRefund.All error: %s", err.Error())
	} else {
		if len(items) > 0 {
			childId = items[0].ID
		}
		t.Logf("items = %s", jsonx.ToPrettyJson(items))
	}
}

func TestOrderRefundService_Create(t *testing.T) {
	t.Run("getOrderId", getOrderId)
	req := CreateOrderRefundRequest{
		Amount:     1,
		Reason:     "product is lost",
		RefundedBy: 0,
		MetaData:   nil,
		LineItems:  nil,
	}
	item, err := wooClient.Services.OrderRefund.Create(mainId, req)
	if err != nil {
		t.Fatalf("wooClient.Services.OrderRefund.Create error: %s", err.Error())
	} else {
		assert.Equal(t, 100.00, item.Amount, "refund amount")
		childId = item.ID
	}
}

func TestOrderRefundService_One(t *testing.T) {
	t.Run("TestOrderRefundService_All", TestOrderRefundService_All)
	item, err := wooClient.Services.OrderRefund.One(mainId, childId, 2)
	if err != nil {
		t.Errorf("wooClient.Services.OrderRefund.One(%d, %d) error: %s", orderId, noteId, err.Error())
	} else {
		assert.Equal(t, childId, item.ID, "note id")
	}
}

func TestOrderRefundService_CreateDelete(t *testing.T) {
	t.Run("getOrderId", getOrderId)
	req := CreateOrderRefundRequest{
		Amount:     100,
		Reason:     "product is lost",
		RefundedBy: 0,
		MetaData:   nil,
		LineItems:  nil,
	}
	item, err := wooClient.Services.OrderRefund.Create(mainId, req)
	if err != nil {
		t.Fatalf("wooClient.Services.OrderRefund.Create error: %s", err.Error())
	} else {
		assert.Equal(t, 100.00, item.Amount, "refund amount")
		noteId = item.ID
	}

	// Delete
	_, err = wooClient.Services.OrderNote.Delete(mainId, childId, true)
	if err != nil {
		t.Fatalf("wooClient.Services.OrderRefund.Delete(%d, %d, %v) error: %s", mainId, childId, true, err.Error())
	} else {
		_, err = wooClient.Services.OrderNote.One(mainId, childId)
		if !errors.Is(err, ErrNotFound) {
			t.Fatalf("wooClient.Services.OrderRefund.Delete(%d, %d, %v) failed", mainId, childId, true)
		}
	}
}
