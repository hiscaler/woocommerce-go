package product

type Term struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	MenuOrder   int    `json:"menu_order"`
	Count       int    `json:"count"`
}
