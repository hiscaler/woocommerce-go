package entity

type ProductDimension struct {
	Length string `json:"length"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

// Product product properties
type Product struct {
	ID                int                       `json:"id"`
	Name              string                    `json:"name"`
	Slug              string                    `json:"slug"`
	Permalink         string                    `json:"permalink"`
	DateCreated       string                    `json:"date_created"`
	DateCreatedGMT    string                    `json:"date_created_gmt"`
	DateModified      string                    `json:"date_modified"`
	DateModifiedGMT   string                    `json:"date_modified_gmt"`
	Type              string                    `json:"type"`
	Status            string                    `json:"status"`
	Featured          bool                      `json:"featured"`
	CatalogVisibility string                    `json:"catalog_visibility"`
	Description       string                    `json:"description"`
	ShortDescription  string                    `json:"short_description"`
	SKU               string                    `json:"sku"`
	Price             string                    `json:"price"`
	RegularPrice      string                    `json:"regular_price"`
	SalePrice         string                    `json:"sale_price"`
	DateOnSaleFrom    string                    `json:"date_on_sale_from"`
	DateOnSaleFromGMT string                    `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      string                    `json:"date_on_sale_to"`
	DateOnSaleToGMT   string                    `json:"date_on_sale_to_gmt"`
	PriceHtml         string                    `json:"price_html"`
	OnSale            bool                      `json:"on_sale"`
	Purchasable       bool                      `json:"purchasable"`
	TotalSales        int                       `json:"total_sales"`
	Virtual           bool                      `json:"virtual"`
	Downloadable      bool                      `json:"downloadable"`
	Downloads         []ProductDownload         `json:"downloads"`
	DownloadLimit     int                       `json:"download_limit"`
	DownloadExpiry    int                       `json:"download_expiry"`
	ExternalUrl       string                    `json:"external_url"`
	ButtonText        string                    `json:"button_text"`
	TaxStatus         string                    `json:"tax_status"`
	TaxClass          string                    `json:"tax_class"`
	ManageStock       bool                      `json:"manage_stock"`
	StockQuantity     int                       `json:"stock_quantity"`
	StockStatus       string                    `json:"stock_status"`
	Backorders        string                    `json:"backorders"`
	BackordersAllowed bool                      `json:"backorders_allowed"`
	Backordered       bool                      `json:"backordered"`
	SoldIndividually  bool                      `json:"sold_individually"`
	Weight            string                    `json:"weight"`
	Dimensions        *ProductDimension         `json:"dimensions"`
	ShippingRequired  bool                      `json:"shipping_required"`
	ShippingTaxable   bool                      `json:"shipping_taxable"`
	ShippingClass     string                    `json:"shipping_class"`
	ShippingClassId   int                       `json:"shipping_class_id"`
	ReviewsAllowed    bool                      `json:"reviews_allowed"`
	AverageRating     string                    `json:"average_rating"`
	RatingCount       int                       `json:"rating_count"`
	RelatedIds        []int                     `json:"related_ids"`
	UpsellIds         []int                     `json:"upsell_ids"`
	CrossSellIds      []int                     `json:"cross_sell_ids"`
	ParentId          int                       `json:"parent_id"`
	PurchaseNote      string                    `json:"purchase_note"`
	Categories        []ProductCategory         `json:"categories"`
	Tags              []ProductTag              `json:"tags"`
	Images            []ProductImage            `json:"images"`
	Attributes        []ProductAttributeItem    `json:"attributes"`
	DefaultAttributes []ProductDefaultAttribute `json:"default_attributes"`
	Variations        []int                     `json:"variations"`
	GroupedProducts   []int                     `json:"grouped_products"`
	MenuOrder         int                       `json:"menu_order"`
	MetaData          []Meta                    `json:"meta_data"`
}

// ProductAttributeItem product attribute properties
type ProductAttributeItem struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Visible   bool     `json:"visible"`
	Variation bool     `json:"variation"`
	Options   []string `json:"options"`
}
