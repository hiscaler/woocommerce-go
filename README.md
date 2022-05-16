WooCommerce SDK for golang
==========================

## 文档

### Rest API Docs

https://woocommerce.github.io/woocommerce-rest-api-docs/#introduction

## 接口

### 商品

- Products(params ProductsQueryParams) (items []product.Product, isLastPage bool, err error)      // 商品列表
- Product(id int) (item product.Product, err error)                                               // 商品详情

### 商品标签

- Tags(params TagsQueryParams) (items []product.Tag, isLastPage bool, err error)                  // 商品标签列表
- Tag(id int) (item product.Tag, err error)                                                       // 商品标签
- CreateTag(req CreateTagRequest) (item product.Tag, err error)                                   // 新增商品标签
- UpdateTag(id int, req UpdateTagRequest) (item product.Tag, err error)                           // 更新商品标签
- DeleteTag(id int) (item product.Tag, err error)                                                 // 删除商品标签
- BatchTags(req CUDTagsRequest) (res BatchTagsResult, err error)                                  // 批量操作商品标签

### 商品分类

- Categories(params CategoriesQueryParams) (items []product.Category, isLastPage bool, err error) // 商品分类列表
- Category(id int) (item product.Category, err error)                                             // 商品分类
- CreateCategory(req CreateCategoryRequest) (item product.Category, err error)                    // 新增商品分类
- UpdateCategory(id int, req UpdateCategoryRequest) (item product.Category, err error)            // 更新商品分类
- DeleteCategory(id int) (item product.Category, err error)                                       // 删除商品标签
- BatchCategories(req CUDCategoriesRequest) (res BatchCategoriesResult, err error)                // 批量操作商品分类