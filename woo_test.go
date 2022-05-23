package woocommerce

import (
	"fmt"
	"github.com/hiscaler/woocommerce-go/config"
	jsoniter "github.com/json-iterator/go"
	"os"
	"testing"
)

var wooClient *WooCommerce

var orderId, noteId int
var mainId, childId int

func getOrderId(t *testing.T) {
	t.Log("Execute getOrderId test")
	items, _, _, _, err := wooClient.Services.Order.All(OrdersQueryParams{})
	if err != nil || len(items) == 0 {
		t.FailNow()
	}
	orderId = items[0].ID
	mainId = items[0].ID
}

func TestMain(m *testing.M) {
	b, err := os.ReadFile("./config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	wooClient = NewClient(c)
	m.Run()
}
