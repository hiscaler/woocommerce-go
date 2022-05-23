package woocommerce

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderNoteService_All(t *testing.T) {
	t.Run("getOrderId", getOrderId)
	params := OrderNotesQueryParams{}
	items, _, _, _, err := wooClient.Services.OrderNote.All(orderId, params)
	if err != nil {
		t.Errorf("wooClient.Services.OrderNote.All error: %s", err.Error())
	} else {
		if len(items) > 0 {
			noteId = items[0].ID
		}
		t.Logf("items = %s", jsonx.ToPrettyJson(items))
	}
}

func TestOrderNoteService_Create(t *testing.T) {
	note := gofakeit.Address().Address
	req := CreateOrderNoteRequest{
		Note: note,
	}
	item, err := wooClient.Services.OrderNote.Create(orderId, req)
	if err != nil {
		t.Fatalf("wooClient.Services.OrderNote.Create error: %s", err.Error())
	} else {
		assert.Equal(t, note, item.Note, "note")
		noteId = item.ID
	}
}

func TestOrderNoteService_One(t *testing.T) {
	t.Run("TestOrderNoteService_All", TestOrderNoteService_All)
	item, err := wooClient.Services.OrderNote.One(orderId, noteId)
	if err != nil {
		t.Errorf("wooClient.Services.OrderNote.One(%d, %d) error: %s", orderId, noteId, err.Error())
	} else {
		assert.Equal(t, noteId, item.ID, "note id")
	}
}

func TestOrderNoteService_CreateDelete(t *testing.T) {
	t.Run("getOrderId", getOrderId)
	note := gofakeit.Address().Address
	req := CreateOrderNoteRequest{
		Note: note,
	}
	item, err := wooClient.Services.OrderNote.Create(orderId, req)
	if err != nil {
		t.Fatalf("wooClient.Services.OrderNote.Create error: %s", err.Error())
	} else {
		assert.Equal(t, note, item.Note, "note")
		noteId = item.ID
	}

	// Delete
	_, err = wooClient.Services.OrderNote.Delete(orderId, noteId, true)
	if err != nil {
		t.Fatalf("wooClient.Services.OrderNote.Delete(%d, %d, %v) error: %s", orderId, noteId, true, err.Error())
	} else {
		_, err = wooClient.Services.OrderNote.One(orderId, noteId)
		if !errors.Is(err, ErrNotFound) {
			t.Fatalf("wooClient.Services.OrderNote.Delete(%d, %d, %v) failed", orderId, noteId, true)
		}
	}
}
