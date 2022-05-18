package entity

// https://woocommerce.github.io/woocommerce-rest-api-docs/?php#retrieve-customer-downloads

type CustomerDownloadFile struct {
	Name string `json:"name"`
	File string `json:"file"`
}
