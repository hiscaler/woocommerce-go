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

**Note**: If the error type is ErrNotFound, it indicates that the corresponding data is not found. If the error type is other error, an error may occur in the call. So you should judge the results to further process your business logic.

## Services

### Coupons

Service Name: wooClient.Service.Coupon

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Customers

Service Name: wooClient.Service.Customer

Methods:

- All
- Batch
- Create
- Delete
- Downloads
- One
- Update

### Order

Service Name: wooClient.Service.Order

Methods:

- All
- Create
- Delete
- One
- Update

### Order Notes

Service Name: wooClient.Service.OrderNote

Methods:

- All
- Create
- Delete
- One

### Refunds

Service Name: wooClient.Service.OrderRefund

Methods:

- All
- Create
- Delete
- One

### Products

Service Name: wooClient.Service.Product

Methods:

- All
- Create
- Delete
- One
- Update

### Product Variations

Service Name: wooClient.Service.ProductVariation

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Product Attributes

Service Name: wooClient.Service.ProductAttribute

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Product Attribute Terms

Service Name: wooClient.Service.ProductAttributeTerm

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Product Categories

Service Name: wooClient.Service.ProductCategory

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Product Shipping Classes

Service Name: wooClient.Service.ProductShippingClass

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Product Tags

Service Name: wooClient.Service.ProductTag

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Product Reviews

Service Name: wooClient.Service.ProductReview

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Report

Service Name: wooClient.Service.Report

Methods:

- All
- CouponTotals
- CustomerTotals
- OrderTotals
- ProductTotals
- ReviewTotals
- SalesReports
- TopSellerReports

### Tax Rates

Service Name: wooClient.Service.TaxRate

Methods:

- All
- Batch
- Create
- Delete
- One
- Update

### Tax Classes

Service Name: wooClient.Service.TaxClass

Methods:

- All
- Create
- Delete

### Settings

- Groups

Service Name: wooClient.Service.Group

Methods:

### Setting Options

- All
- One
- Update

### Payment Gateways

Service Name: wooClient.Service.PaymentGateway

Methods:

- All
- One
- Update

### Shipping Zones

Service Name: wooClient.Service.ShippingZone

Methods:

- All
- Create
- Delete
- One
- Update

### Shipping Zone Locations

Service Name: wooClient.Service.ShippingZoneLocation

Methods:

- All
- Update

### Shipping Zone Methods

Service Name: wooClient.Service.ShippingZoneMethod

Methods:

- All
- Delete
- Include
- One
- Update

### Shipping Methods

Service Name: wooClient.Service.ShippingMethod

Methods:

- All
- One

### System Status

Service Name: wooClient.Service.SystemStatus

Methods:

- All

### System Status Tools

Service Name: wooClient.Service.SystemStatusTool

Methods:

- All
- One
- Run

### Data

Service Name: wooClient.Service.Data

Methods:

- All
- Continent
- Continents
- Countries
- Country
- Currencies
- Currency
- CurrentCurrency


## Contributing
If you have any questions or suggestions, you can:

1. [Report an issue](https://github.com/hiscaler/woocommerce-go/issues/new)
2. Fork it and pull a request

Thanks.