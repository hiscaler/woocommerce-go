WooCommerce SDK for golang
==========================

## Docs

### Rest API Docs

https://woocommerce.github.io/woocommerce-rest-api-docs/#introduction

## Usage

### Step 1. Create a new client

```go
// Read you config
b, err := os.ReadFile("./config/config_test.json")
if err != nil {
    panic(fmt.Sprintf("Read config error: %s", err.Error()))
}
var c config.Config
err = jsoniter.Unmarshal(b, &c)
if err != nil {
    panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
}

wooClient := NewClient(c)
```

Now you get a wooCommerce client object, If you want operate data, please refer second step. 

### Step 2. Call special service method
```go
// Product Query
params := ProductsQueryParams{}
wooClient.Services.Product.All(params)
```

10 records are returned by default.

In most cases, you can filter by condition, example:

```go
params.SKU = "NO123"
wooClient.Services.Product.All(params)
```

The first ten eligible records are returned in this case

And you can retrieve one data use One() method.

```go
product, err := wooClient.Services.Product.One(1)
```

**Note**: If the error type is ErrNotFound, it indicates that the corresponding data is not found. If the error type is other error, an error may occur in the call.  So you should judge the results to further process your business logic.


## Services

### Coupons

- All
- One
- Create
- Delete
- Batch

## APIs（Will be discarded）

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
- Customer(id int) (item customer.Customer, err error)                                           // Retrieve a customer
- CreateCustomer(req CreateCustomerRequest) (item customer.Customer, err error)                                     // Create a customer
- UpdateCustomer(req UpdateCustomerRequest) (item customer.Customer, err error)                  // Update a customer

### Order Notes

- OrderNotes(orderId int, params OrderNotesQueryParams) (items []order.Note, isLastPage bool, err error) // List all order notes
- OrderNote(orderId, noteId int) (item order.Note, err error)                                            // Retrieve an order note
- CreateOrderNote(orderId int, req CreateOrderNoteRequest) (item order.Note, err error)                  // Create an order note
- DeleteOrderNote(orderId, noteId int, force bool) (item order.Note, err error)                          // Delete an order note

### Tax Classes

- TaxClasses() (items []TaxClass, err error)                           // List all tax classes
- CreateTaxClass(req CreateTaxClassRequest) (item TaxClass, err error) // Create a tax class
- DeleteTaxClass(slug string) (item TaxClass, err error)               // Delete a tax class