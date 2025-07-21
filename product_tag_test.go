package woocommerce

import (
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
)

func TestProductTagService_All(t *testing.T) {
	params := ProductTagsQueryParams{}
	items, _, _, _, err := wooClient.Services.ProductTag.All(params)
	if err != nil {
		t.Errorf("wooClient.Services.ProductTag.All: %s", err.Error())
	} else {
		t.Logf("Items: %s", jsonx.ToPrettyJson(items))
	}
}

func TestProductTagService_One(t *testing.T) {
	item, err := wooClient.Services.ProductTag.One(51)
	if err != nil {
		t.Errorf("wooClient.Services.ProductTag.One: %s", err.Error())
	} else {
		assert.Equal(t, 51, item.ID, "one")
	}
}

func TestProductTagService_CreateUpdateDelete(t *testing.T) {
	name := gofakeit.BeerName()
	req := CreateProductTagRequest{
		Name: name,
	}
	item, err := wooClient.Services.ProductTag.Create(req)
	if err != nil {
		t.Fatalf("wooClient.Services.ProductTag.Create: %s", err.Error())
	}
	assert.Equal(t, name, item.Name, "product tag name")
	tagId := item.ID

	// Update
	newName := gofakeit.BeerName()
	updateReq := UpdateProductTagRequest{
		Name: newName,
	}
	item, err = wooClient.Services.ProductTag.Update(tagId, updateReq)
	if err != nil {
		t.Fatalf("wooClient.Services.ProductTag.Update: %s", err.Error())
	}
	assert.Equal(t, newName, item.Name, "product tag name")

	// Delete
	_, err = wooClient.Services.ProductTag.Delete(tagId, true)
	if err != nil {
		t.Fatalf("wooClient.Services.ProductTag.Delete: %s", err.Error())
	}

	// Check is exists
	_, err = wooClient.Services.ProductTag.One(tagId)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("%d is not deleted, error: %s", tagId, err.Error())
	}
}

func TestProductTagService_Batch(t *testing.T) {
	n := 3
	createRequests := make([]BatchProductTagsCreateItem, n)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		req := BatchProductTagsCreateItem{
			Name: gofakeit.Word(),
		}
		createRequests[i] = req
		names[i] = req.Name
	}
	batchReq := BatchProductTagsRequest{
		Create: createRequests,
	}
	result, err := wooClient.Services.ProductTag.Batch(batchReq)
	if err != nil {
		t.Fatalf("wooClient.Services.ProductTag.Batch() error: %s", err.Error())
	}
	assert.Equal(t, n, len(result.Create), "Batch create return len")
	returnNames := make([]string, 0)
	for _, d := range result.Create {
		returnNames = append(returnNames, d.Name)
	}
	assert.Equal(t, names, returnNames, "check names is equal")
}
