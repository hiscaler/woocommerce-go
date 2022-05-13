package product

import "time"

type Image struct {
	Id              int       `json:"id"`
	DateCreated     time.Time `json:"date_created"`
	DateCreatedGMT  time.Time `json:"date_created_gmt"`
	DateModified    time.Time `json:"date_modified"`
	DateModifiedGMT time.Time `json:"date_modified_gmt"`
	Src             string    `json:"src"`
	Name            string    `json:"name"`
	Alt             string    `json:"alt"`
}
