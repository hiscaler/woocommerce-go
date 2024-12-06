package woocommerce

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dashboard-bg/woocommerce-go/entity"
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
)

func getCustomerId(t *testing.T) {
	t.Log("Execute getCustomerId test")
	params := CustomersQueryParams{}
	params.Page = 1
	params.PerPage = 1
	items, _, _, _, err := wooClient.Services.Customer.All(params)
	if err != nil || len(items) == 0 {
		t.FailNow()
	}
	if len(items) == 0 {
		t.Fatalf("getCustomerId not found one customer")
	}
	mainId = items[0].ID
}

func TestCustomerService_All(t *testing.T) {
	params := CustomersQueryParams{}
	_, _, _, _, err := wooClient.Services.Customer.All(params)
	if err != nil {
		t.Errorf("wooClient.Services.Customer.All error: %s", err.Error())
	}
}

func TestCustomerService_One(t *testing.T) {
	t.Run("getCustomerId", getCustomerId)
	item, err := wooClient.Services.Customer.One(mainId)
	if err != nil {
		t.Fatalf("wooClient.Services.Customer.One error: %s", err.Error())
	}
	assert.Equal(t, mainId, item.ID, "customer id")
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
	var oldItem, newItem entity.Customer
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
	oldItem, err = wooClient.Services.Customer.Create(req)
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
	_, err = wooClient.Services.Customer.Delete(oldItem.ID, customerDeleteParams{Force: true})
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

func TestCustomerService_Downloads(t *testing.T) {
	// todo
	items, err := wooClient.Services.Customer.Downloads(0)
	if err != nil {
		t.Fatalf("wooClient.Services.Customer.Downloads() error: %s", err.Error())
	} else {
		t.Logf("items = %s", jsonx.ToPrettyJson(items))
	}
}
