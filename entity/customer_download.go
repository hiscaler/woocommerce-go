package entity

// https://woocommerce.github.io/woocommerce-rest-api-docs/?php#retrieve-customer-downloads

// CustomerDownload customer download properties
type CustomerDownload struct {
	DownloadId        string               `json:"download_id"`
	DownloadURL       string               `json:"download_url"`
	ProductId         string               `json:"product_id"`
	ProductName       string               `json:"product_name"`
	DownloadName      string               `json:"download_name"`
	OrderId           int                  `json:"order_id"`
	OrderKey          string               `json:"order_key"`
	DownloadRemaining string               `json:"download_remaining"`
	AccessExpires     string               `json:"access_expires"`
	AccessExpiresGMT  string               `json:"access_expires_gmt"`
	File              CustomerDownloadFile `json:"file"`
}
