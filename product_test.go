package woocommerce

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestProductService_All(t *testing.T) {
	params := ProductsQueryParams{}
	products, _, err := wooClient.Services.Product.All(params)
	if err != nil {
		t.Errorf("wooClient.Services.Product.All: %s", err.Error())
	} else {
		t.Logf("Products: %s", jsonx.ToPrettyJson(products))
	}
}

func TestProductService_One(t *testing.T) {
	product, err := wooClient.Services.Product.One(625)
	if err != nil {
		t.Errorf("wooClient.Services.Product.One: %s", err.Error())
	} else {
		t.Logf("Product: %s", jsonx.ToPrettyJson(product))
	}
}

func TestProductService_Variations(t *testing.T) {
	params := ProductVariationsQueryParams{}
	variations, _, err := wooClient.Services.Product.Variations(1, params)
	if err != nil {
		t.Errorf("wooClient.Services.Product.Variations: %s", err.Error())
	} else {
		t.Logf("Variations: %s", jsonx.ToPrettyJson(variations))
	}
}
