package entity

// ProductReview product review properties
type ProductReview struct {
	ID             int    `json:"id"`
	DateCreated    string `json:"date_created"`
	DateCreatedGMT string `json:"date_created_gmt"`
	ProductId      int    `json:"product_id"`
	Status         string `json:"status"`
	Reviewer       string `json:"reviewer"`
	ReviewerEmail  string `json:"reviewer_email"`
	Review         string `json:"review"`
	Rating         int    `json:"rating"`
	Verified       bool   `json:"verified"`
}
