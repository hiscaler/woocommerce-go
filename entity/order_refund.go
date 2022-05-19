package entity

// OrderRefund order refund properties
type OrderRefund struct {
	ID              int                   `json:"id"`
	DateCreated     string                `json:"date_created"`
	DateCreatedGMT  string                `json:"date_created_gmt"`
	Amount          string                `json:"amount"`
	Reason          string                `json:"reason"`
	RefundedBy      int                   `json:"refunded_by"`
	RefundedPayment bool                  `json:"refunded_payment"`
	MetaData        []MetaData            `json:"meta_data"`
	LineItems       []OrderRefundLineItem `json:"line_items"`
	APIRefund       bool                  `json:"api_refund"`
}

// OrderRefundLineItem order refund line item properties
type OrderRefundLineItem struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ProductId   int        `json:"product_id"`
	VariationId int        `json:"variation_id"`
	Quantity    int        `json:"quantity"`
	TaxClass    int        `json:"tax_class"`
	SubTotal    string     `json:"subtotal"`
	SubTotalTax string     `json:"subtotal_tax"`
	Total       string     `json:"total"`
	TotalTax    string     `json:"total_tax"`
	Taxes       []Taxes    `json:"taxes"`
	MetaData    []MetaData `json:"meta_data"`
	SKU         string     `json:"sku"`
	Price       float64    `json:"price"`
	RefundTotal float64    `json:"refund_total"`
}
