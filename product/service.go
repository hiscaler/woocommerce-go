package product

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/product"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Products(params ProductsQueryParams) (items []product.Product, isLastPage bool, err error)      // 商品列表
	Product(id int) (item product.Product, err error)                                               // 商品详情
	Tags(params TagsQueryParams) (items []product.Tag, isLastPage bool, err error)                  // 商品标签列表
	Tag(id int) (item product.Tag, err error)                                                       // 商品标签
	CreateTag(req UpsertTagRequest) (item product.Tag, err error)                                   // 新增商品标签
	UpdateTag(id int, req UpsertTagRequest) (item product.Tag, err error)                           // 更新商品标签
	DeleteTag(id int) (item product.Tag, err error)                                                 // 删除商品标签
	BatchTags(req CUDTagsRequest) (res BatchTagsResult, err error)                                  // 批量操作商品标签
	Categories(params CategoriesQueryParams) (items []product.Category, isLastPage bool, err error) // 商品分类列表
	Category(id int) (item product.Category, err error)                                             // 商品分类
	CreateCategory(req UpsertCategoryRequest) (item product.Category, err error)                    // 新增商品分类
	UpdateCategory(id int, req UpsertCategoryRequest) (item product.Category, err error)            // 更新商品分类
	DeleteCategory(id int) (item product.Category, err error)                                       // 删除商品标签
	BatchCategories(req CUDCategoriesRequest) (res BatchCategoriesResult, err error)                // 批量操作商品分类
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
