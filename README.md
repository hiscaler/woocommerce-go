WooCommerce SDK for golang
==========================

## Docs

### Rest API Docs

https://woocommerce.github.io/woocommerce-rest-api-docs/#introduction

## APIs

### Products

- Products(params ProductsQueryParams) (items []product.Product, isLastPage bool, err error)      // List all products
- Product(id int) (item product.Product, err error)                                               // Retrieve a product

categories

### Product Tags

- Tags(params TagsQueryParams) (items []product.Tag, isLastPage bool, err error)                  // List all product tags
- Tag(id int) (item product.Tag, err error)                                                       // Retrieve a product tag
- CreateTag(req CreateTagRequest) (item product.Tag, err error)                                   // Create a product tag
- UpdateTag(id int, req UpdateTagRequest) (item product.Tag, err error)                           // Update a product tag
- DeleteTag(id int) (item product.Tag, err error)                                                 // Delete a product tag
- BatchTags(req CUDTagsRequest) (res BatchTagsResult, err error)                                  // Batch update product tags

### Product Categories

- Categories(params CategoriesQueryParams) (items []product.Category, isLastPage bool, err error) // List all product categories
- Category(id int) (item product.Category, err error)                                             // Retrieve a product category
- CreateCategory(req CreateCategoryRequest) (item product.Category, err error)                    // Create a product category
- UpdateCategory(id int, req UpdateCategoryRequest) (item product.Category, err error)            // Update a product category
- DeleteCategory(id int) (item product.Category, err error)                                       // Delete a product category
- BatchCategories(req CUDCategoriesRequest) (res BatchCategoriesResult, err error)                // Batch update product categories

### Orders

- Orders(params OrdersQueryParams) (items []order.Order, isLastPage bool, err error) // List all orders
- Order(id int) (item order.Order, err error)                                        // Retrieve an order

### System Status

- SystemStatus() (item systemstatus.SystemStatus, err error) // List all system status items

### Settings

- Groups() (items []Group, err error) // List all settings groups

### Tax Rates

- TaxRates(params TaxRatesQueryParams) (items []TaxRate, isLastPage bool, err error) // List all tax rates

### Customers

- Customers(params CustomersQueryParams) (items []customer.Customer, isLastPage bool, err error) // List all customers
- Customer(id int) (item customer.Customer, err error)

### Order Notes

- OrderNotes(orderId int, params OrderNotesQueryParams) (items []order.Note, isLastPage bool, err error) // List all order notes
- OrderNote(orderId, noteId int) (item order.Note, err error)                                            // Retrieve an order note
- CreateOrderNote(orderId int, req CreateOrderNoteRequest) (item order.Note, err error)                  // Create an order note
- DeleteOrderNote(orderId, noteId int, force bool) (item order.Note, err error)                          // Delete an order note