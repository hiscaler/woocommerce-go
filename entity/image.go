package entity

// ProductImage product iamge properties
type ProductImage struct {
	ID              int    `json:"id"`
	DateCreated     string `json:"date_created"`
	DateCreatedGMT  string `json:"date_created_gmt"`
	DateModified    string `json:"date_modified"`
	DateModifiedGMT string `json:"date_modified_gmt"`
	Src             string `json:"src"`
	Name            string `json:"name"`
	Alt             string `json:"alt"`
}
