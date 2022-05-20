package woocommerce

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productService service

// Products

type ProductsQueryParams struct {
	queryParams
	Search        string   `url:"search,omitempty"`
	After         string   `url:"after,omitempty"`
	Before        string   `url:"before,omitempty"`
	Exclude       []string `url:"exclude,omitempty"`
	Include       []string `url:"include,omitempty"`
	Parent        []string `url:"parent,omitempty"`
	ParentExclude []string `url:"parent_exclude,omitempty"`
	Slug          string   `url:"slug,omitempty"`
	Status        string   `url:"status,omitempty"`
	Type          string   `url:"type,omitempty"`
	SKU           string   `url:"sku,omitempty"`
	Featured      bool     `url:"featured,omitempty"`
	Category      string   `url:"category,omitempty"`
	Tag           string   `url:"tag,omitempty"`
	ShippingClass string   `url:"shipping_class,omitempty"`
	Attribute     string   `url:"attribute,omitempty"`
	AttributeTerm string   `url:"attribute_term,omitempty"`
	TaxClass      string   `url:"tax_class,omitempty"`
	OnSale        bool     `url:"on_sale,omitempty"`
	MinPrice      string   `url:"min_price,omitempty"`
	MaxPrice      string   `url:"max_price,omitempty"`
	StockStatus   string   `url:"stock_status,omitempty"`
}

func (m ProductsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "title", "slug", "price", "popularity", "rating").Error("无效的排序字段"))),
		validation.Field(&m.Status, validation.When(m.Status != "", validation.In("any", "draft", "pending", "private", "publish").Error("无效的状态"))),
		validation.Field(&m.Type, validation.When(m.Type != "", validation.In("simple", "grouped", "external", "variable").Error("无效的类型"))),
	)
}

// All List all products
func (s productService) All(params ProductsQueryParams) (items []entity.Product, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	urlValues, _ := query.Values(params)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get("/products")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &items); err == nil {
			isLastPage = len(items) < params.PerPage
		}
	}
	return
}

// One Retrieve a product
func (s productService) One(id int) (item entity.Product, err error) {
	var res entity.Product
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			item = res
		}
	}
	return
}

// Create

type CreateProductRequest struct {
	Name              string                           `json:"name,omitempty"`
	Slug              string                           `json:"slug,omitempty"`
	Type              string                           `json:"type,omitempty"`
	Status            string                           `json:"status,omitempty"`
	Featured          bool                             `json:"featured,omitempty"`
	CatalogVisibility string                           `json:"catalog_visibility,omitempty"`
	Description       string                           `json:"description,omitempty"`
	ShortDescription  string                           `json:"short_description,omitempty"`
	SKU               string                           `json:"sku,omitempty"`
	RegularPrice      string                           `json:"regular_price,omitempty"`
	SalePrice         string                           `json:"sale_price,omitempty"`
	DateOnSaleFrom    string                           `json:"date_on_sale_from,omitempty"`
	DateOnSaleFromGMT string                           `json:"date_on_sale_from_gmt,omitempty"`
	DateOnSaleTo      string                           `json:"date_on_sale_to,omitempty"`
	DateOnSaleToGMT   string                           `json:"date_on_sale_to_gmt,omitempty"`
	Virtual           bool                             `json:"virtual,omitempty"`
	Downloadable      bool                             `json:"downloadable,omitempty"`
	Downloads         []entity.ProductDownload         `json:"downloads,omitempty"`
	DownloadLimit     int                              `json:"download_limit,omitempty"`
	DownloadExpiry    int                              `json:"download_expiry,omitempty"`
	ExternalUrl       string                           `json:"external_url,omitempty"`
	ButtonText        string                           `json:"button_text,omitempty"`
	TaxStatus         string                           `json:"tax_status,omitempty"`
	TaxClass          string                           `json:"tax_class,omitempty"`
	ManageStock       bool                             `json:"manage_stock,omitempty"`
	StockQuantity     int                              `json:"stock_quantity,omitempty"`
	StockStatus       string                           `json:"stock_status,omitempty"`
	Backorders        string                           `json:"backorders,omitempty"`
	SoldIndividually  bool                             `json:"sold_individually,omitempty"`
	Weight            string                           `json:"weight,omitempty"`
	Dimensions        *entity.ProductDimension         `json:"dimensions,omitempty"`
	ShippingClass     string                           `json:"shipping_class,omitempty"`
	ReviewsAllowed    bool                             `json:"reviews_allowed,omitempty"`
	UpsellIds         []int                            `json:"upsell_ids,omitempty"`
	CrossSellIds      []int                            `json:"cross_sell_ids,omitempty"`
	ParentId          int                              `json:"parent_id,omitempty"`
	PurchaseNote      string                           `json:"purchase_note,omitempty"`
	Categories        []entity.ProductCategory         `json:"categories,omitempty"`
	Tags              []entity.ProductTag              `json:"tags,omitempty"`
	Images            []entity.ProductImage            `json:"images,omitempty"`
	Attributes        []entity.ProductAttribute        `json:"attributes,omitempty"`
	DefaultAttributes []entity.ProductDefaultAttribute `json:"default_attributes,omitempty"`
	GroupedProducts   []int                            `json:"grouped_products,omitempty"`
	MenuOrder         int                              `json:"menu_order,omitempty"`
	MetaData          []entity.Meta                    `json:"meta_data,omitempty"`
}

func (m CreateProductRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required.Error("商品名称不能为空")),
		validation.Field(&m.Type, validation.When(m.Type != "", validation.In("simple", "grouped", "external ", "variable").Error("无效的类型"))),
		validation.Field(&m.Status, validation.When(m.Status != "", validation.In("draft", "pending", "private", "publish").Error("无效的状态"))),
		validation.Field(&m.CatalogVisibility, validation.When(m.CatalogVisibility != "", validation.In("visible", "catalog", "search", "hidden").Error("无效的目录可见性"))),
		validation.Field(&m.TaxStatus, validation.When(m.TaxStatus != "", validation.In("taxable", "shipping ", "none").Error("无效的税务状态"))),
		validation.Field(&m.StockStatus, validation.When(m.StockStatus != "", validation.In("instock", "outofstock ", "onbackorder").Error("无效的库存状态"))),
		validation.Field(&m.Backorders, validation.When(m.Backorders != "", validation.In("yes", "no ", "notify").Error("无效的缺货订单状态"))),
	)
}

// Create create a product
func (s productService) Create(req CreateProductRequest) (item entity.Product, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Update

type UpdateProductRequest = CreateProductRequest

func (s productService) Update(id int, req UpdateProductRequest) (item entity.Product, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Put(fmt.Sprintf("/products/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete delete a product
func (s productService) Delete(id int, force bool) (item entity.Product, err error) {
	resp, err := s.httpClient.R().SetBody(map[string]bool{"force": force}).Delete(fmt.Sprintf("/products/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
