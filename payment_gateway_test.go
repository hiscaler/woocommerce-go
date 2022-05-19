package woocommerce

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/hiscaler/woocommerce-go/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var paymentGatewayId string

func TestPaymentGatewayService_All(t *testing.T) {
	items, err := wooClient.Services.PaymentGateway.All()
	if err != nil {
		t.Fatalf("wooClient.Services.PaymentGateway.All error: %s", err.Error())
	}
	if len(items) > 0 {
		paymentGatewayId = items[0].ID
	}
}

func TestPaymentGatewayService_One(t *testing.T) {
	t.Run("TestPaymentGatewayService_All", TestPaymentGatewayService_All)
	item, err := wooClient.Services.PaymentGateway.One(paymentGatewayId)
	if err != nil {
		t.Errorf("wooClient.Services.Coupon.PaymentGateway error: %s", err.Error())
	} else {
		assert.Equal(t, paymentGatewayId, item.ID, "payment gateway id")
	}
}

func TestPaymentGatewayService_Update(t *testing.T) {
	t.Run("TestPaymentGatewayService_All", TestPaymentGatewayService_All)

	var oldItem, newItem entity.PaymentGateway
	var err error
	oldItem, err = wooClient.Services.PaymentGateway.One(paymentGatewayId)
	if err != nil {
		t.Fatalf("wooClient.Services.PaymentGateway.One error: %s", err.Error())
	}

	req := UpdatePaymentGatewayRequest{}
	newItem, err = wooClient.Services.PaymentGateway.Update(paymentGatewayId, req)
	if err != nil {
		t.Fatalf("wooClient.Services.PaymentGateway.Update error: %s", err.Error())
	}
	assert.Equal(t, oldItem, newItem, "all no change")

	// Change title
	req.Title = gofakeit.RandomString([]string{"A", "B", "C", "D", "E", "F", "G"})
	newItem, err = wooClient.Services.PaymentGateway.Update(paymentGatewayId, req)
	if err != nil {
		t.Fatalf("wooClient.Services.PaymentGateway.Update error: %s", err.Error())
	}
	assert.Equal(t, req.Title, newItem.Title, "title")
}
