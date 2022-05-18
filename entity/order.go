package entity

type LineItem struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ProductId   int        `json:"product_id"`
	VariationId int        `json:"variation_id"`
	Quantity    int        `json:"quantity"`
	TaxClass    string     `json:"tax_class"`
	SubTotal    string     `json:"subtotal"`
	SubTotalTax string     `json:"subtotal_tax"`
	Total       string     `json:"total"`
	TotalTax    string     `json:"total_tax"`
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
	Total     string     `json:"total"`
	TotalTax  string     `json:"total_tax"`
	Taxes     []Taxes    `json:"taxes"`
	MetaData  []MetaData `json:"meta_data"`
}

type CouponLine struct {
	ID          int        `json:"id"`
	Code        string     `json:"code"`
	Discount    string     `json:"discount"`
	DiscountTax string     `json:"discount_tax"`
	MetaData    []MetaData `json:"meta_data"`
}

type Refund struct {
	ID     int    `json:"id"`
	Reason string `json:"reason"`
	Total  string `json:"total"`
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
	DiscountTotal      string         `json:"discount_total"`
	DiscountTax        string         `json:"discount_tax"`
	ShippingTotal      string         `json:"shipping_total"`
	ShippingTax        string         `json:"shipping_tax"`
	CartTax            string         `json:"cart_tax"`
	Total              string         `json:"total"`
	TotalTax           string         `json:"total_tax"`
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
