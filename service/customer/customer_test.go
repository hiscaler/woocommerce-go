package customer

import (
	"fmt"
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/config"
	"github.com/hiscaler/woocommerce-go/entity"
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
	customers, _, err := wooService.Customers(params)
	if err != nil {
		t.Errorf("wooService.Customers error: %s", err.Error())
	} else {
		t.Logf("customers = %#v", customers)
	}
}

func TestService_Customer(t *testing.T) {
	customer, err := wooService.Customer(4)
	if err != nil {
		t.Errorf("wooService.Customer error: %s", err.Error())
	} else {
		t.Logf("customer = %#v", customer)
	}
}

func TestService_CreateCustomer(t *testing.T) {
	req := CreateCustomerRequest{
		Email:     "lisi2@example.com",
		FirstName: "zhang",
		LastName:  "san",
		Username:  "lisi2",
		Password:  "123",
		MetaData:  nil,
		Billing: &entity.Billing{
			FirstName: "zhangsan",
			LastName:  "wife",
			Company:   "xx company",
			Address1:  "China HN",
			Address2:  "",
			City:      "CS",
			State:     "",
			Postcode:  "410000",
			Country:   "China",
			Email:     "john@example.com",
			Phone:     "1",
		},
	}
	customer, err := wooService.CreateCustomer(req)
	if err != nil {
		t.Errorf("wooService.CreateCustomer error: %s", err.Error())
	} else {
		t.Logf("customer = %#v", customer)
	}
}
