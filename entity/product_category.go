package entity

// ProductCategory product category properties
type ProductCategory struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Parent      int           `json:"parent"`
	Description string        `json:"description"`
	Display     string        `json:"display"`
	Image       *ProductImage `json:"image,omitempty"`
	MenuOrder   int           `json:"menu_order"`
	Count       int           `json:"count"`
}
