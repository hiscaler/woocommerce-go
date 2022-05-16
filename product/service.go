package product

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/product"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Products(params ProductsQueryParams) (items []product.Product, isLastPage bool, err error) // 商品列表
	Product(id int) (item product.Product, err error)                                          // 商品详情
	Tags(params TagsQueryParams) (items []product.Tag, isLastPage bool, err error)             // 商品标签
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
