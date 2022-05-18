package entity

// ProductTag product tag properties
type ProductTag struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}
