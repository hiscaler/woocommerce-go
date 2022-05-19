package woocommerce

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductCategoryService_All(t *testing.T) {
	params := ProductCategoriesQueryParams{}
	items, _, err := wooClient.Services.ProductCategory.All(params)
	if err != nil {
		t.Errorf("wooClient.Services.ProductCategory.All: %s", err.Error())
	} else {
		t.Logf("Items: %s", jsonx.ToPrettyJson(items))
	}
}

func TestProductCategoryService_One(t *testing.T) {
	item, err := wooClient.Services.ProductCategory.One(15)
	if err != nil {
		t.Errorf("wooClient.Services.ProductCategory.One: %s", err.Error())
	} else {
		assert.Equal(t, 15, item.ID, "one")
	}
}

func TestProductCategoryService_CreateUpdateDelete(t *testing.T) {
	name := gofakeit.BeerName()
	req := CreateProductCategoryRequest{
		Name: name,
	}
	item, err := wooClient.Services.ProductCategory.Create(req)
	if err != nil {
		t.Fatalf("wooClient.Services.ProductCategory.Create: %s", err.Error())
	}
	assert.Equal(t, name, item.Name, "product category name")
	categoryId := item.ID

	// Update
	newName := gofakeit.BeerName()
	updateReq := UpdateProductCategoryRequest{
		Name: newName,
	}
	item, err = wooClient.Services.ProductCategory.Update(categoryId, updateReq)
	if err != nil {
		t.Fatalf("wooClient.Services.ProductCategory.Update: %s", err.Error())
	}
	assert.Equal(t, newName, item.Name, "product category name")

	// Delete
	_, err = wooClient.Services.ProductCategory.Delete(categoryId, true)
	if err != nil {
		t.Fatalf("wooClient.Services.ProductCategory.Delete: %s", err.Error())
	}

	// Check is exists
	_, err = wooClient.Services.ProductCategory.One(categoryId)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("%d is not deleted, error: %s", categoryId, err.Error())
	}
}

func TestProductCategoryService_Batch(t *testing.T) {
	n := 3
	createRequests := make([]BatchProductCategoriesCreateItem, n)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		req := BatchProductCategoriesCreateItem{
			Name:        gofakeit.Word(),
			Description: gofakeit.Address().Address,
		}
		createRequests[i] = req
		names[i] = req.Name
	}
	batchReq := BatchProductCategoriesRequest{
		Create: createRequests,
	}
	result, err := wooClient.Services.ProductCategory.Batch(batchReq)
	if err != nil {
		t.Fatalf("wooClient.Services.ProductCategory.Batch() error: %s", err.Error())
	}
	assert.Equal(t, n, len(result.Create), "Batch create return len")
	returnNames := make([]string, 0)
	for _, d := range result.Create {
		returnNames = append(returnNames, d.Name)
	}
	assert.Equal(t, names, returnNames, "check names is equal")
}
