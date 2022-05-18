package woocommerce

import (
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/gox/randx"
	"github.com/hiscaler/woocommerce-go/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerService_All(t *testing.T) {
	params := CustomersQueryParams{}
	items, _, err := wooClient.Services.Customer.All(params)
	if err != nil {
		t.Errorf("wooClient.Services.Customer.All error: %s", err.Error())
	} else {
		t.Logf("items = %#v", jsonx.ToPrettyJson(items))
	}
}

func TestCustomerService_One(t *testing.T) {
	item, err := wooClient.Services.Customer.One(4)
	if err != nil {
		t.Errorf("wooClient.Services.Customer.One error: %s", err.Error())
	} else {
		t.Logf("item = %s", jsonx.ToPrettyJson(item))
	}
}

func TestCustomerService_Create(t *testing.T) {
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
	customer, err := wooClient.Services.Customer.Create(req)
	if err != nil {
		t.Errorf("wooClient.Services.Customer.Create error: %s", err.Error())
	} else {
		t.Logf("customer = %#v", customer)
	}
}

func TestCustomerService_Update(t *testing.T) {
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
	customer, err := wooClient.Services.Customer.Create(req)
	if err != nil {
		t.Errorf("wooClient.Services.Customer.Create error: %s", err.Error())
	} else {
		newUsername := randx.Letter(6, false)
		updateReq := UpdateCustomerRequest{
			Email: newUsername + "@example.com",
			Billing: &entity.Billing{
				FirstName: "billfn-" + newUsername,
				LastName:  "billln-" + newUsername,
			},
		}
		customer, err = wooClient.Services.Customer.Update(customer.ID, updateReq)
		if err != nil {
			t.Errorf("wooClient.Services.Customer.Update error: %s", err.Error())
		} else {
			t.Log(jsonx.ToPrettyJson(customer))
			assert.Equal(t, newUsername+"@example.com", customer.Email, "email")
			assert.Equal(t, "billfn-"+newUsername, customer.Billing.FirstName, "billing first name")
			assert.Equal(t, "billln-"+newUsername, customer.Billing.LastName, "billing last name")
		}
	}
}

func TestCustomerService_Delete(t *testing.T) {
	username := randx.Letter(6, false)
	req := CreateCustomerRequest{
		Email:     username + "@example.com",
		FirstName: "zhang",
		LastName:  "san",
		Username:  username,
		Password:  "123",
		MetaData:  nil,
		Billing: &entity.Billing{
			FirstName: "billfn-" + username,
			LastName:  "billln-" + username,
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
	customer, err := wooClient.Services.Customer.Create(req)
	if err != nil {
		t.Errorf("wooClient.Services.Customer.Create error: %s", err.Error())
	} else {
		_, err = wooClient.Services.Customer.Delete(customer.ID)
		if err != nil {
			// todo 501: Customers do not support trashing.
			t.Errorf("wooClient.Services.Customer.Delete error: %s", err.Error())
		} else {
			_, err = wooClient.Services.Customer.One(customer.ID)
			if err == nil {
				t.Errorf("wooClient.Services.Customer.One(%d) is exists, not deleted.", customer.ID)
			}
		}
	}
}
