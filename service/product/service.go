package product

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity/product"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Products(params ProductsQueryParams) (items []product.Product, isLastPage bool, err error)      // List all products
	Product(id int) (item product.Product, err error)                                               // Retrieve a product
	Tags(params TagsQueryParams) (items []product.Tag, isLastPage bool, err error)                  // List all product tags
	Tag(id int) (item product.Tag, err error)                                                       // Retrieve a product tag
	CreateTag(req CreateTagRequest) (item product.Tag, err error)                                   // Create a product tag
	UpdateTag(id int, req UpdateTagRequest) (item product.Tag, err error)                           // Update a product tag
	DeleteTag(id int) (item product.Tag, err error)                                                 // Delete a product tag
	BatchTags(req CUDTagsRequest) (res BatchTagsResult, err error)                                  // Batch update product tags
	Categories(params CategoriesQueryParams) (items []product.Category, isLastPage bool, err error) // List all product categories
	Category(id int) (item product.Category, err error)                                             // Retrieve a product category
	CreateCategory(req CreateCategoryRequest) (item product.Category, err error)                    // Create a product category
	UpdateCategory(id int, req UpdateCategoryRequest) (item product.Category, err error)            // Update a product category
	DeleteCategory(id int) (item product.Category, err error)                                       // Delete a product category
	BatchCategories(req CUDCategoriesRequest) (res BatchCategoriesResult, err error)                // Batch update product categories
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
