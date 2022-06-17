package entity

type LineItem struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ProductId   int        `json:"product_id"`
	VariationId int        `json:"variation_id"`
	Quantity    int        `json:"quantity"`
	TaxClass    string     `json:"tax_class"`
	SubTotal    float64    `json:"subtotal"`
	SubTotalTax float64    `json:"subtotal_tax"`
	Total       float64    `json:"total"`
	TotalTax    float64    `json:"total_tax"`
	Taxes       []Taxes    `json:"taxes"`
	MetaData    []MetaData `json:"meta_data"`
	SKU         string     `json:"sku"`
	Price       float64    `json:"price"`
	ParentName  string     `json:"parent_name"`
}

type FeeLine struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	TaxClass  string     `json:"tax_class"`
	TaxStatus string     `json:"tax_status"`
	Total     float64    `json:"total"`
	TotalTax  float64    `json:"total_tax"`
	Taxes     []Taxes    `json:"taxes"`
	MetaData  []MetaData `json:"meta_data"`
}

type CouponLine struct {
	ID          int        `json:"id"`
	Code        string     `json:"code"`
	Discount    float64    `json:"discount"`
	DiscountTax float64    `json:"discount_tax"`
	MetaData    []MetaData `json:"meta_data"`
}

type Refund struct {
	ID     int     `json:"id"`
	Reason string  `json:"reason"`
	Total  float64 `json:"total"`
}

// Order order properties
type Order struct {
	ID                 int            `json:"id"`
	ParentId           int            `json:"parent_id"`
	Number             string         `json:"number"`
	OrderKey           string         `json:"order_key"`
	CreatedVia         string         `json:"created_via"`
	Version            string         `json:"version"`
	Status             string         `json:"status"`
	Currency           string         `json:"currency"`
	CurrencySymbol     string         `json:"currency_symbol"`
	DateCreated        string         `json:"date_created"`
	DateCreatedGMT     string         `json:"date_created_gmt"`
	DateModified       string         `json:"date_modified"`
	DateModifiedGMT    string         `json:"date_modified_gmt"`
	DiscountTotal      float64        `json:"discount_total"`
	DiscountTax        float64        `json:"discount_tax"`
	ShippingTotal      float64        `json:"shipping_total"`
	ShippingTax        float64        `json:"shipping_tax"`
	CartTax            float64        `json:"cart_tax"`
	Total              float64        `json:"total"`
	TotalTax           float64        `json:"total_tax"`
	PricesIncludeTax   bool           `json:"prices_include_tax"`
	CustomerId         int            `json:"customer_id"`
	CustomerIpAddress  string         `json:"customer_ip_address"`
	CustomerUserAgent  string         `json:"customer_user_agent"`
	CustomerNote       string         `json:"customer_note"`
	Billing            Billing        `json:"billing"`
	Shipping           Shipping       `json:"shipping"`
	PaymentMethod      string         `json:"payment_method"`
	PaymentMethodTitle string         `json:"payment_method_title"`
	TransactionId      string         `json:"transaction_id"`
	DatePaid           string         `json:"date_paid"`
	DatePaidGMT        string         `json:"date_paid_gmt"`
	DateCompleted      string         `json:"date_completed"`
	DateCompletedGMT   string         `json:"date_completed_gmt"`
	CartHash           string         `json:"cart_hash"`
	MetaData           []MetaData     `json:"meta_data"`
	LineItems          []LineItem     `json:"line_items"`
	TaxLines           []TaxLine      `json:"tax_lines"`
	ShippingLines      []ShippingLine `json:"shipping_lines"`
	FeeLines           []FeeLine      `json:"fee_lines"`
	CouponLines        []CouponLine   `json:"coupon_lines"`
	Refunds            []Refund       `json:"refunds"`
	SetPaid            bool           `json:"set_paid"`
}
