package order

import (
	"github.com/hiscaler/woocommerce-go/entity"
	"time"
)

type LineItem struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	ProductId   int               `json:"product_id"`
	VariationId int               `json:"variation_id"`
	Quantity    int               `json:"quantity"`
	TaxClass    string            `json:"tax_class"`
	SubTotal    string            `json:"subtotal"`
	SubTotalTax string            `json:"subtotal_tax"`
	Total       string            `json:"total"`
	TotalTax    string            `json:"total_tax"`
	Taxes       []entity.Taxes    `json:"taxes"`
	MetaData    []entity.MetaData `json:"meta_data"`
	SKU         string            `json:"sku"`
	Price       float64           `json:"price"`
	ParentName  string            `json:"parent_name"`
}

type FeeLine struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	TaxClass  string            `json:"tax_class"`
	TaxStatus string            `json:"tax_status"`
	Total     string            `json:"total"`
	TotalTax  string            `json:"total_tax"`
	Taxes     []entity.Taxes    `json:"taxes"`
	MetaData  []entity.MetaData `json:"meta_data"`
}

type CouponLine struct {
	Id          int               `json:"id"`
	Code        string            `json:"code"`
	Discount    string            `json:"discount"`
	DiscountTax string            `json:"discount_tax"`
	MetaData    []entity.MetaData `json:"meta_data"`
}

type Refund struct {
	Id     int    `json:"id"`
	Reason string `json:"reason"`
	Total  string `json:"total"`
}

type Order struct {
	Id                 int                   `json:"id"`
	ParentId           int                   `json:"parent_id"`
	Number             string                `json:"number"`
	OrderKey           string                `json:"order_key"`
	CreatedVia         string                `json:"created_via"`
	Version            string                `json:"version"`
	Status             string                `json:"status"`
	Currency           string                `json:"currency"`
	CurrencySymbol     string                `json:"currency_symbol"`
	DateCreated        time.Time             `json:"date_created"`
	DateCreatedGMT     time.Time             `json:"date_created_gmt"`
	DateModified       time.Time             `json:"date_modified"`
	DateModifiedGMT    time.Time             `json:"date_modified_gmt"`
	DiscountTotal      string                `json:"discount_total"`
	DiscountTax        string                `json:"discount_tax"`
	ShippingTotal      string                `json:"shipping_total"`
	ShippingTax        string                `json:"shipping_tax"`
	CartTax            string                `json:"cart_tax"`
	Total              string                `json:"total"`
	TotalTax           string                `json:"total_tax"`
	PricesIncludeTax   bool                  `json:"prices_include_tax"`
	CustomerId         int                   `json:"customer_id"`
	CustomerIpAddress  string                `json:"customer_ip_address"`
	CustomerUserAgent  string                `json:"customer_user_agent"`
	CustomerNote       string                `json:"customer_note"`
	Billing            entity.Billing        `json:"billing"`
	Shipping           entity.Shipping       `json:"shipping"`
	PaymentMethod      string                `json:"payment_method"`
	PaymentMethodTitle string                `json:"payment_method_title"`
	TransactionId      string                `json:"transaction_id"`
	DatePaid           time.Time             `json:"date_paid"`
	DatePaidGMT        time.Time             `json:"date_paid_gmt"`
	DateCompleted      time.Time             `json:"date_completed"`
	DateCompletedGMT   time.Time             `json:"date_completed_gmt"`
	CartHash           string                `json:"cart_hash"`
	MetaData           []entity.MetaData     `json:"meta_data"`
	LineItems          []LineItem            `json:"line_items"`
	TaxLines           []entity.TaxLine      `json:"tax_lines"`
	ShippingLines      []entity.ShippingLine `json:"shipping_lines"`
	FeeLines           []FeeLine             `json:"fee_lines"`
	CouponLines        []CouponLine          `json:"coupon_lines"`
	Refunds            []Refund              `json:"refunds"`
	SetPaid            bool                  `json:"set_paid"`
}
