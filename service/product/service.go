package product

import (
	"github.com/hiscaler/woocommerce-go"
	"github.com/hiscaler/woocommerce-go/entity"
)

type service struct {
	woo *woocommerce.WooCommerce
}

type Service interface {
	Products(params ProductsQueryParams) (items []entity.Product, isLastPage bool, err error)             // List all products
	Product(id int) (item entity.Product, err error)                                                      // Retrieve a product
	Tags(params TagsQueryParams) (items []entity.ProductTag, isLastPage bool, err error)                  // List all product tags
	Tag(id int) (item entity.ProductTag, err error)                                                       // Retrieve a product tag
	CreateTag(req CreateTagRequest) (item entity.ProductTag, err error)                                   // Create a product tag
	UpdateTag(id int, req UpdateTagRequest) (item entity.ProductTag, err error)                           // Update a product tag
	DeleteTag(id int) (item entity.ProductTag, err error)                                                 // Delete a product tag
	BatchTags(req CUDTagsRequest) (res BatchTagsResult, err error)                                        // Batch update product tags
	Categories(params CategoriesQueryParams) (items []entity.ProductCategory, isLastPage bool, err error) // List all product categories
	Category(id int) (item entity.ProductCategory, err error)                                             // Retrieve a product category
	CreateCategory(req CreateCategoryRequest) (item entity.ProductCategory, err error)                    // Create a product category
	UpdateCategory(id int, req UpdateCategoryRequest) (item entity.ProductCategory, err error)            // Update a product category
	DeleteCategory(id int) (item entity.ProductCategory, err error)                                       // Delete a product category
	BatchCategories(req CUDCategoriesRequest) (res BatchCategoriesResult, err error)                      // Batch update product categories
}

func NewService(client *woocommerce.WooCommerce) Service {
	return service{client}
}
