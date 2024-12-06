package woocommerce

import (
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

type productVariationService service

// Product variations

type ProductVariationsQueryParams struct {
	queryParams
	Search        string  `url:"search,omitempty"`
	After         string  `url:"after,omitempty"`
	Before        string  `url:"before,omitempty"`
	Exclude       []int   `url:"exclude,omitempty"`
	Include       []int   `url:"include,omitempty"`
	Parent        []int   `url:"parent,omitempty"`
	ParentExclude []int   `url:"parent_exclude,omitempty"`
	Slug          string  `url:"slug,omitempty"`
	Status        string  `url:"status,omitempty"`
	SKU           string  `url:"sku,omitempty"`
	TaxClass      string  `url:"tax_class,omitempty"`
	OnSale        string  `url:"on_sale,omitempty"`
	MinPrice      float64 `url:"min_price,omitempty"`
	MaxPrice      float64 `url:"max_price,omitempty"`
	StockStatus   string  `url:"stock_status,omitempty"`
}

func (m ProductVariationsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Before, validation.When(m.Before != "", validation.By(func(value interface{}) error {
			dateStr, _ := value.(string)
			return IsValidateTime(dateStr)
		}))),
		validation.Field(&m.After, validation.When(m.After != "", validation.By(func(value interface{}) error {
			dateStr, _ := value.(string)
			return IsValidateTime(dateStr)
		}))),
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "title", "include", "date", "slug").Error("invalid sort field"))),
		validation.Field(&m.Status, validation.When(m.Status != "", validation.In("any", "draft", "pending", "private", "publish").Error("Invalid status value"))),
		validation.Field(&m.TaxClass, validation.When(m.TaxClass != "", validation.In("standard", "reduced-rate", "zero-rate").Error("Invalid tax class"))),
		validation.Field(&m.StockStatus, validation.When(m.StockStatus != "", validation.In("instock", "outofstock", "onbackorder").Error("Invalid stock status"))),
		validation.Field(&m.MinPrice, validation.Min(0.0)),
		validation.Field(&m.MaxPrice, validation.Min(m.MinPrice)),
	)
}

// All List all product variations
func (s productVariationService) All(productId int, params ProductVariationsQueryParams) (items []entity.ProductVariation, total, totalPages int, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	params.TidyVars()
	params.After = ToISOTimeString(params.After, false, true)
	params.Before = ToISOTimeString(params.Before, true, false)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(toValues(params)).Get(fmt.Sprintf("/products/%d/variations", productId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
		total, totalPages, isLastPage = parseResponseTotal(params.Page, resp)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// One retrieve a product variation
func (s productVariationService) One(productId, variationId int) (item entity.ProductVariation, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/products/%d/variations/%d", productId, variationId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Create

type CreateProductVariationRequest struct {
	Description    string                             `json:"description,omitempty"`
	SKU            string                             `json:"sku,omitempty"`
	RegularPrice   float64                            `json:"regular_price,string,omitempty"`
	SalePrice      float64                            `json:"sale_price,string,omitempty"`
	Status         string                             `json:"status,omitempty"`
	Virtual        bool                               `json:"virtual,omitempty"`
	Downloadable   bool                               `json:"downloadable,omitempty"`
	Downloads      []entity.ProductDownload           `json:"downloads,omitempty"`
	DownloadLimit  int                                `json:"download_limit,omitempty"`
	DownloadExpiry int                                `json:"download_expiry,omitempty"`
	TaxStatus      string                             `json:"tax_status,omitempty"`
	TaxClass       string                             `json:"tax_class,omitempty"`
	ManageStock    bool                               `json:"manage_stock,omitempty"`
	StockQuantity  int                                `json:"stock_quantity,omitempty"`
	StockStatus    string                             `json:"stock_status,omitempty"`
	Backorders     string                             `json:"backorders,omitempty"`
	Weight         float64                            `json:"weight,string,omitempty"`
	Dimension      *entity.ProductDimension           `json:"dimensions,omitempty"`
	ShippingClass  string                             `json:"shipping_class,omitempty"`
	Image          *entity.ProductImage               `json:"image,omitempty"`
	Attributes     []entity.ProductVariationAttribute `json:"attributes,omitempty"`
	MenuOrder      int                                `json:"menu_order,omitempty"`
	MetaData       []entity.Meta                      `json:"meta_data,omitempty"`
}

func (m CreateProductVariationRequest) Validate() error {
	return nil
}

func (s productVariationService) Create(productId int, req CreateProductVariationRequest) (item entity.ProductVariation, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
		SetBody(req).
		Post(fmt.Sprintf("/products/%d/variations", productId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Update

type UpdateProductVariationRequest = CreateProductVariationRequest

func (s productVariationService) Update(productId int, req UpdateProductVariationRequest) (item entity.ProductVariation, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
		SetBody(req).
		Put(fmt.Sprintf("/products/%d/variations", productId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Delete

func (s productVariationService) Delete(productId, variationId int, force bool) (item entity.ProductVariation, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]bool{"force": force}).
		Delete(fmt.Sprintf("/products/%d/variations/%d", productId, variationId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	} else {
		err = ErrorWrap(resp.StatusCode(), "")
	}
	return
}

// Batch Update

type BatchProductVariationsCreateItem = CreateProductVariationRequest
type BatchProductVariationsUpdateItem struct {
	ID int `json:"id"`
	CreateProductVariationRequest
}

type BatchProductVariationsRequest struct {
	Create []BatchProductVariationsCreateItem `json:"create,omitempty"`
	Update []BatchProductVariationsUpdateItem `json:"update,omitempty"`
	Delete []int                              `json:"delete,omitempty"`
}

func (m BatchProductVariationsRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("invalid request data")
	}
	return nil
}

type BatchProductVariationsResult struct {
	Create []entity.ProductVariation `json:"create"`
	Update []entity.ProductVariation `json:"update"`
	Delete []entity.ProductVariation `json:"delete"`
}

func (s productVariationService) Batch(req BatchProductVariationsRequest) (res BatchProductVariationsResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/products/variations/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
