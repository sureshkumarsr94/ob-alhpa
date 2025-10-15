package stock_dto

import (
	distributorDto "infopack.co.in/offybox/app/dto/distributor"
	masterDto "infopack.co.in/offybox/app/dto/master"
	productDto "infopack.co.in/offybox/app/dto/product"
	coreModel "infopack.co.in/offybox/app/models/core"
)

type StockRequest struct {
	StockID       string `json:"stock_id"`
	WarehouseID   string `json:"warehouse_id" validate:"required"`
	DistributorID string `json:"distributor_id" validate:"required"`
	UserID        string `json:"user_id"`
	ProductID     string `json:"product_id" validate:"required"`
	Quantity      int    `json:"quantity" validate:"required"`
	Description   string `json:"description"`
	BatchNo       string `json:"batch_no" validate:"required"`
	Status        string `json:"status" validate:"required"`
}

type StockResponse struct {
	StockID       string                             `json:"stock_id"`
	WarehouseID   string                             `json:"warehouse_id"`
	Warehouse     masterDto.WarehouseResponse        `json:"warehouse"`
	DistributorID string                             `json:"distributor_id"`
	Distributor   distributorDto.DistributorResponse `json:"distributor"`
	UserID        string                             `json:"user_id"`
	User          coreModel.User                     `json:"user"`
	ProductID     string                             `json:"product_id"`
	Product       productDto.ProductResponse         `json:"product"`
	Quantity      int                                `json:"quantity"`
	Description   string                             `json:"description"`
	BatchNo       string                             `json:"batch_no"`
	Status        string                             `json:"status"`
}
