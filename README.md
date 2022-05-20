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

- All
- Batch
- Create
- Delete
- One
- Update

### Customers

- All
- Batch
- Create
- Delete
- Downloads
- One
- Update

### Order

- All
- Create
- Delete
- One
- Update

### Order Notes

- All
- Create
- Delete
- One

### Refunds

- All
- Create
- Delete
- One

### Products

- All
- Create
- Delete
- One
- Update

### Product Variations

- All
- Batch
- Create
- Delete
- One
- Update

### Product Attributes

- All
- Batch
- Create
- Delete
- One
- Update

### Product Attribute Terms

- All
- Batch
- Create
- Delete
- One
- Update

### Product Categories

- All
- Batch
- Create
- Delete
- One
- Update

### Product Shipping Classes

- All
- Batch
- Create
- Delete
- One
- Update

### Product Tags

- All
- Batch
- Create
- Delete
- One
- Update

### Product Reviews

- All
- Batch
- Create
- Delete
- One
- Update

### Report
- All
- CouponTotals
- CustomerTotals
- OrderTotals
- ProductTotals
- ReviewTotals
- SalesReports
- TopSellerReports

### Tax Rates

- All
- Batch
- Create
- Delete
- One
- Update

### Tax Classes

- All
- Create
- Delete

### Settings

- Groups

### Setting Options

- All
- One
- Update

### Payment Gateways

- All
- One
- Update

### Shipping Zones

- All
- Create
- Delete
- One
- Update

### Shipping Zone Locations

- All
- Update

### Shipping Zone Methods

- All
- Delete
- Include
- One
- Update

### Shipping Methods

- All
- One

### System Status

- All

### System Status Tools

- All
- One
- Run

### Data

- All
- Continent
- Continents
- Countries
- Country
- Currencies
- Currency
- CurrentCurrency