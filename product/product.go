package product

import (
	"github.com/hiscaler/woocommerce-go/entity/product"
	jsoniter "github.com/json-iterator/go"
)

type ProductsQueryParams struct {
	Context       string   `json:"context,omitempty"`
	Search        string   `json:"search,omitempty"`
	After         string   `json:"after,omitempty"`
	Before        string   `json:"before,omitempty"`
	Exclude       []string `json:"exclude,omitempty"`
	Include       []string `json:"include,omitempty"`
	Offset        int      `json:"offset,omitempty"`
	Order         string   `json:"order,omitempty"`
	OrderBy       string   `json:"orderby,omitempty"`
	Parent        []string `json:"parent,omitempty"`
	ParentExclude []string `json:"parent_exclude,omitempty"`
	Slug          string   `json:"slug,omitempty"`
	Status        string   `json:"status,omitempty"`
	Type          string   `json:"type,omitempty"`
	SKU           string   `json:"sku,omitempty"`
	Featured      bool     `json:"featured,omitempty"`
	Category      string   `json:"category,omitempty"`
	Tag           string   `json:"tag,omitempty"`
	ShippingClass string   `json:"shipping_class,omitempty"`
	Attribute     string   `json:"attribute,omitempty"`
	AttributeTerm string   `json:"attribute_term,omitempty"`
	TaxClass      string   `json:"tax_class,omitempty"`
	OnSale        bool     `json:"on_sale,omitempty"`
	MinPrice      string   `json:"min_price,omitempty"`
	MaxPrice      string   `json:"max_price,omitempty"`
	StockStatus   string   `json:"stock_status,omitempty"`
}

func (m ProductsQueryParams) Validate() error {
	return nil
}

func (s service) Products(params ProductsQueryParams) (items []product.Product, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	res := struct {
		Data []product.Product `json:"data"`
	}{}
	qp := make(map[string]string, 0)
	resp, err := s.woo.Client.R().
		SetQueryParams(qp).
		Get("/products")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			items = res.Data
		}
	}
	return
}
