package entity

// OrderNote order note properties
type OrderNote struct {
	ID             int    `json:"id"`
	Author         string `json:"author"`
	DateCreated    string `json:"date_created"`
	DateCreatedGMT string `json:"date_created_gmt"`
	Note           string `json:"note"`
	CustomerNote   bool   `json:"customer_note"`
	AddedByUser    bool   `json:"added_by_user"`
}
