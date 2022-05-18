package woocommerce

import (
	"github.com/hiscaler/gox/jsonx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductTagService_All(t *testing.T) {
	params := ProductTagsQueryParams{}
	items, _, err := wooClient.Services.ProductTag.All(params)
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
