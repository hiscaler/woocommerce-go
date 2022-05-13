package product

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Parent      int    `json:"parent"`
	Description string `json:"description"`
	Display     string `json:"display"`
	Image       *Image `json:"image,omitempty"`
	MenuOrder   int    `json:"menu_order"`
	Count       int    `json:"count"`
}
