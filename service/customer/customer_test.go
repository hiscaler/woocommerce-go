package customer

import (
	"fmt"
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/config"
	jsoniter "github.com/json-iterator/go"
	"os"
	"testing"
)

var wooInstance *woocommerce.WooCommerce
var wooService Service

func TestMain(m *testing.M) {
	b, err := os.ReadFile("../../config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	wooInstance = woocommerce.NewClient(c)
	wooService = NewService(wooInstance)
	m.Run()
}

func TestService_Customers(t *testing.T) {
	params := CustomersQueryParams{}
	orders, _, err := wooService.Customers(params)
	if err != nil {
		t.Errorf("wooService.Customers error: %s", err.Error())
	} else {
		t.Logf("orders = %#v", orders)
	}
}

func TestService_Customer(t *testing.T) {
	order, err := wooService.Customer(1)
	if err != nil {
		t.Errorf("wooService.Customer error: %s", err.Error())
	} else {
		t.Logf("order = %#v", order)
	}
}
