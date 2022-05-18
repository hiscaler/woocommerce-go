package woocommerce

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/woocommerce-go/entity"
	customer2 "github.com/hiscaler/woocommerce-go/entity/customer"
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
	gofakeit.Seed(0)
	address := gofakeit.Address()
	req := CreateCustomerRequest{
		Email:     gofakeit.Email(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Username:  gofakeit.Username(),
		Password:  gofakeit.Password(true, true, true, false, false, 10),
		MetaData:  nil,
		Billing: &entity.Billing{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Company:   gofakeit.Company(),
			Address1:  address.Address,
			Address2:  "",
			City:      address.City,
			State:     address.State,
			Postcode:  address.Zip,
			Country:   address.Country,
			Email:     gofakeit.Email(),
			Phone:     gofakeit.Phone(),
		},
	}
	item, err := wooClient.Services.Customer.Create(req)
	if err != nil {
		t.Errorf("wooClient.Services.Customer.Create error: %s", err.Error())
	} else {
		t.Logf("item = %#v", item)
	}
}

func TestCustomerService_CreateUpdateDelete(t *testing.T) {
	gofakeit.Seed(0)
	// Create
	var oldItem, newItem customer2.Customer
	var err error
	address := gofakeit.Address()
	req := CreateCustomerRequest{
		Email:     gofakeit.Email(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Username:  gofakeit.Username(),
		Password:  gofakeit.Password(true, true, true, false, false, 10),
		MetaData:  nil,
		Billing: &entity.Billing{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Company:   gofakeit.Company(),
			Address1:  address.Address,
			Address2:  "",
			City:      address.City,
			State:     address.State,
			Postcode:  address.Zip,
			Country:   address.Country,
			Email:     gofakeit.Email(),
			Phone:     gofakeit.Phone(),
		},
	}
	newItem, err = wooClient.Services.Customer.Create(req)
	if err != nil {
		t.Fatalf("wooClient.Services.Customer.Create error: %s", err.Error())
	}

	// Update
	afterData := struct {
		email            string
		billingFirstName string
		billingLastName  string
	}{
		email:            gofakeit.Email(),
		billingFirstName: gofakeit.FirstName(),
		billingLastName:  gofakeit.LastName(),
	}
	updateReq := UpdateCustomerRequest{
		Email: afterData.email,
		Billing: &entity.Billing{
			FirstName: afterData.billingFirstName,
			LastName:  afterData.billingLastName,
		},
	}
	newItem, err = wooClient.Services.Customer.Update(oldItem.ID, updateReq)
	if err != nil {
		t.Fatalf("wooClient.Services.Customer.Update error: %s", err.Error())
	} else {
		assert.Equal(t, afterData.email, newItem.Email, "email")
		assert.Equal(t, afterData.billingFirstName, newItem.Billing.FirstName, "billing first name")
		assert.Equal(t, afterData.billingLastName, newItem.Billing.LastName, "billing last name")
	}

	// Delete
	_, err = wooClient.Services.Customer.Delete(oldItem.ID)
	if err != nil {
		t.Fatalf("wooClient.Services.Customer.Delete(%d) error: %s", oldItem.ID, err.Error())
	}

	// Query check is exists
	_, err = wooClient.Services.Customer.One(oldItem.ID)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("wooClient.Services.Customer.Delete(%d) failed", oldItem.ID)
	}
}

func TestCustomerService_Batch(t *testing.T) {
	n := 3
	createRequests := make([]BatchCreateCustomerRequest, n)
	emails := make([]string, n)
	for i := 0; i < n; i++ {
		req := BatchCreateCustomerRequest{
			Email:     gofakeit.Email(),
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Username:  gofakeit.Username(),
			Password:  gofakeit.Password(true, true, true, false, false, 10),
			Billing:   nil,
			Shipping:  nil,
			MetaData:  nil,
		}
		createRequests[i] = req
		emails[i] = req.Email
	}
	batchReq := BatchCustomerRequest{
		Create: createRequests,
	}
	result, err := wooClient.Services.Customer.Batch(batchReq)
	if err != nil {
		t.Fatalf("wooClient.Services.Customer.Batch() error: %s", err.Error())
	}
	assert.Equal(t, n, len(result.Create), "Batch create return len")
	returnEmails := make([]string, 0)
	for _, d := range result.Create {
		returnEmails = append(returnEmails, d.Email)
	}
	assert.Equal(t, emails, returnEmails, "check emails is equal")
}
