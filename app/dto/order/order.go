package order_dto

import (
	outletDto "infopack.co.in/offybox/app/dto/outlet"
	productDto "infopack.co.in/offybox/app/dto/product"
	userDto "infopack.co.in/offybox/app/dto/user"
)

type OrderRequest struct {
	OrderID         string             `json:"order_id" `
	OutletID        string             `json:"outlet_id" validate:"required"`
	OutletAddressID string             `json:"outlet_address_id" validate:"required"`
	Type            string             `json:"type" validate:"required,oneof=ORDER QUOTATION"`
	Remarks         string             `json:"remarks"`
	Latitude        string             `json:"latitude"`
	Longitude       string             `json:"longitude"`
	OrderStatus     string             `json:"order_status" validate:"required"`
	Item            []OrderItemRequest `json:"item" validate:"required"`
}

type OrderItemRequest struct {
	OrderItemID string  `json:"order_item_id"`
	ProductID   string  `json:"product_id" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	Price       float64 `json:"price"`
	Status      string  `json:"status" validate:"required,oneof=ACTIVE,INACTIVE"`
}

type OrderResponse struct {
	OrderID         string                          `json:"order_id"`
	OutletID        string                          `json:"outlet_id"`
	Outlet          outletDto.OutletResponse        `json:"outlet"`
	Type            string                          `json:"type"`
	ItemCount       string                          `json:"item_count"`
	NetAmount       string                          `json:"net_amount"`
	DiscountAmount  string                          `json:"discount_amount"`
	TaxAmount       string                          `json:"tax_amount"`
	TotalAmount     string                          `json:"total_amount"`
	QuotationDate   string                          `json:"quotation_date"`
	CreatedByUserID string                          `json:"created_by_user_id"`
	CreatedByUser   userDto.UserObject              `json:"created_by_user"`
	Remarks         string                          `json:"remarks"`
	Latitude        string                          `json:"latitude"`
	Longitude       string                          `json:"longitude"`
	OutletAddressID string                          `json:"outlet_address_id"`
	OutletAddress   outletDto.OutletAddressResponse `json:"outlet_address"`
	FileURL         string                          `json:"file_url"`
	OrderStatus     string                          `json:"order_status"`
	Status          string                          `json:"status"`
	Item            []OrderItemResponse             `json:"item"`
}

type OrderItemResponse struct {
	OrderItemID    string                     `json:"order_item_id"`
	ProductID      string                     `json:"product_id"`
	Product        productDto.ProductResponse `json:"product"`
	FreeProductID  string                     `json:"free_product_id"`
	FreeProduct    productDto.ProductResponse `json:"free_product,omitempty"`
	Quantity       string                     `json:"quantity"`
	FreeQuantity   string                     `json:"free_quantity"`
	Price          string                     `json:"price"`
	NetAmount      string                     `json:"net_amount"`
	DiscountAmount string                     `json:"discount_amount"`
	TaxAmount      string                     `json:"tax_amount"`
	TotalAmount    string                     `json:"total_amount"`
	TaxPercentage  string                     `json:"tax_percentage"`
	SchemeID       string                     `json:"scheme_id"`
	OrderID        string                     `json:"order_id"`
	Status         string                     `json:"status"`
}
