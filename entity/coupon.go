package entity

// Coupon coupon properties
type Coupon struct {
	ID                        int      `json:"id"`
	Code                      string   `json:"code"`
	Amount                    float64  `json:"amount"`
	DateCreated               string   `json:"date_created"`
	DateCreatedGMT            string   `json:"date_created_gmt"`
	DateModified              string   `json:"date_modified"`
	DateModifiedGMT           string   `json:"date_modified_gmt"`
	DiscountType              string   `json:"discount_type"`
	Description               string   `json:"description"`
	DateExpires               string   `json:"date_expires"`
	DateExpiresGMT            string   `json:"date_expires_gmt"`
	UsageCount                int      `json:"usage_count"`
	IndividualUse             bool     `json:"individual_use"`
	ProductIDs                []int    `json:"product_ids"`
	ExcludedProductIDs        []int    `json:"excluded_product_ids"`
	UsageLimit                int      `json:"usage_limit"`
	UsageLimitPerUser         int      `json:"usage_limit_per_user"`
	LimitUsageToXItems        int      `json:"limit_usage_to_x_items"`
	FreeShipping              bool     `json:"free_shipping"`
	ProductCategories         []int    `json:"product_categories"`
	ExcludedProductCategories []int    `json:"excluded_product_categories"`
	ExcludeSaleItems          bool     `json:"exclude_sale_items"`
	MinimumAmount             float64  `json:"minimum_amount"`
	MaximumAmount             float64  `json:"maximum_amount"`
	EmailRestrictions         []string `json:"email_restrictions"`
	UsedBy                    []int    `json:"used_by"`
	MetaData                  []Meta   `json:"meta_data"`
}
