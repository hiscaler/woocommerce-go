package woocommerce

const (
	Version       = "1.0.0"
	UserAgent     = "WooCommerce API Client-Golang/" + Version
	HashAlgorithm = "HMAC-SHA256"
)

type WooCommerce struct {
}

func NewClient() *WooCommerce {
	return &WooCommerce{}
}
