package woocommerce

import (
	"github.com/hiscaler/gox/jsonx"
	"testing"
)

func TestProductVariationService_All(t *testing.T) {
	params := ProductVariationsQueryParams{}
	items, _, _, _, err := wooClient.Services.ProductVariation.All(1, params)
	if err != nil {
		t.Errorf("wooClient.Services.ProductVariation.All: %s", err.Error())
	} else {
		t.Logf("items: %s", jsonx.ToPrettyJson(items))
	}
}
