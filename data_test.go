package woocommerce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataService_All(t *testing.T) {
	_, err := wooClient.Services.Data.All()
	assert.Equal(t, nil, err)
}

func TestDataService_Countries(t *testing.T) {
	_, err := wooClient.Services.Data.Countries()
	assert.Equal(t, nil, err)
}

func TestDataService_Currencies(t *testing.T) {
	_, err := wooClient.Services.Data.Currencies()
	assert.Equal(t, nil, err)
}

func TestDataService_Continents(t *testing.T) {
	_, err := wooClient.Services.Data.Continents()
	assert.Equal(t, nil, err)
}

func TestDataService_Continent(t *testing.T) {
	_, err := wooClient.Services.Data.Continent("AF")
	assert.Equal(t, nil, err)
}
