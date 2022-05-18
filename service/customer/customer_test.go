package customer

import (
	"fmt"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/gox/randx"
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/config"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
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

func TestService_UpdateCustomer(t *testing.T) {
	oldUsername := randx.Letter(6, false)
	req := CreateCustomerRequest{
		Email:     oldUsername + "@example.com",
		FirstName: "zhang",
		LastName:  "san",
		Username:  oldUsername,
		Password:  "123",
		MetaData:  nil,
		Billing: &entity.Billing{
			FirstName: "billfn-" + oldUsername,
			LastName:  "billln-" + oldUsername,
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
		newUsername := randx.Letter(6, false)
		updateReq := UpdateCustomerRequest{
			Email: newUsername + "@example.com",
			Billing: &entity.Billing{
				FirstName: "billfn-" + newUsername,
				LastName:  "billln-" + newUsername,
			},
		}
		customer, err = wooService.UpdateCustomer(customer.ID, updateReq)
		if err != nil {
			t.Errorf("wooService.UpdateCustomer error: %s", err.Error())
		} else {
			t.Log(jsonx.ToPrettyJson(customer))
			assert.Equal(t, newUsername+"@example.com", customer.Email, "email")
			assert.Equal(t, "billfn-"+newUsername, customer.Billing.FirstName, "billing first name")
			assert.Equal(t, "billln-"+newUsername, customer.Billing.LastName, "billing last name")
		}
	}
}
