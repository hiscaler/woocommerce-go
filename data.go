package woocommerce

import (
	"fmt"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type dataService service

// All list all data
func (s dataService) All() (items []entity.Data, err error) {
	resp, err := s.httpClient.R().Get("/data")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// Continents list all continents
func (s dataService) Continents() (items []entity.Continent, err error) {
	resp, err := s.httpClient.R().Get("/data/continents")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// Continent retrieve continent data
func (s dataService) Continent(code string) (item entity.Continent, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/data/continents/%s", code))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Countries list all countries
func (s dataService) Countries() (items []entity.Country, err error) {
	resp, err := s.httpClient.R().Get("/data/countries")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// Country retrieve country data
func (s dataService) Country(code string) (item entity.Country, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/data/countries/%s", code))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Currencies list all currencies
func (s dataService) Currencies() (items []entity.Currency, err error) {
	resp, err := s.httpClient.R().Get("/data/currencies")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// Currency retrieve currency data
func (s dataService) Currency(code string) (item entity.Currency, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/data/currencies/%s", code))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// CurrentCurrency retrieve current currency
func (s dataService) CurrentCurrency() (item entity.Currency, err error) {
	resp, err := s.httpClient.R().Get("/data/currencies/current")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
