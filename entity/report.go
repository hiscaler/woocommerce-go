package entity

// Report report properties
type Report struct {
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type SaleReport struct {
	TotalSales     float64  `json:"total_sales"`
	NetSales       float64  `json:"net_sales"`
	AverageSales   string   `json:"average_sales"`
	TotalOrders    int      `json:"total_orders"`
	TotalItems     int      `json:"total_items"`
	TotalTax       float64  `json:"total_tax"`
	TotalShipping  float64  `json:"total_shipping"`
	TotalRefunds   int      `json:"total_refunds"`
	TotalDiscount  int      `json:"total_discount"`
	TotalGroupedBy string   `json:"total_grouped_by"`
	Totals         []string `json:"totals"`
}

// TopSellerReport top sellers report properties
type TopSellerReport struct {
	Title     string `json:"title"`
	ProductId int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type TotalReport struct {
	Slug  string  `json:"slug"`
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

// CouponTotal coupon total properties
type CouponTotal TotalReport

// CustomerTotal customer total properties
type CustomerTotal TotalReport

// OrderTotal order total properties
type OrderTotal TotalReport

// ProductTotal product total properties
type ProductTotal TotalReport

// ReviewTotal review total properties
type ReviewTotal TotalReport
