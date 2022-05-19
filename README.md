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

### Customers

- All
- One
- Create
- Delete
- Batch
- Downloads

### Order

- All
- One

### Order Notes

- All
- One
- Create
- Delete

### Order Refunds

- All
- One
- Create
- Delete

### Products 

- All
- One
- Create
- Update
- Delete

### Product Categories

- All
- One
- Create
- Delete
- Batch

### Product Tags

- All
- One
- Create
- Delete
- Batch
