package trip_dto

import (
	orderDto "infopack.co.in/offybox/app/dto/order"
	outletDto "infopack.co.in/offybox/app/dto/outlet"
	productDto "infopack.co.in/offybox/app/dto/product"
	coreModel "infopack.co.in/offybox/app/models/core"
)

type TripRequest struct {
	TripID              string            `json:"trip_id"`
	UserID              string            `json:"user_id" validate:"required"`
	StartDate           string            `json:"start_date" validate:"required"`
	EndDate             string            `json:"end_date"`
	VehicleNumber       string            `json:"vehicle_number" validate:"required"`
	VehicleType         string            `json:"vehicle_type" validate:"required"`
	VehicleName         string            `json:"vehicle_name" validate:"required"`
	DriverName          string            `json:"driver_name" validate:"required"`
	DriverContactNumber string            `json:"driver_contact_number" validate:"required"`
	DriverProof         string            `json:"driver_proof" validate:"required"`
	StartKM             string            `json:"start_km" validate:"required"`
	EndKM               string            `json:"end_km"`
	LoadedQuantity      string            `json:"loaded_quantity" validate:"required"`
	ReturnedQuantity    string            `json:"returned_quantity"`
	DamagedQuantity     string            `json:"damaged_quantity"`
	Status              string            `json:"status" validate:"required"`
	Items               []TripItemRequest `json:"items"`
}

type TripItemRequest struct {
	TripItemID   string  `json:"trip_item_id"`
	InvoiceID    string  `json:"invoice_id" validate:"required"`
	OutletID     string  `json:"outlet_id" validate:"required"`
	ProductID    string  `json:"product_id" validate:"required"`
	Quantity     float64 `json:"quantity" validate:"required"`
	FreeQuantity float64 `json:"free_quantity" validate:"required"`
	Remarks      string  `json:"remarks"`
	Status       string  `json:"status" validate:"required"`
}

type TripEndRequest struct {
	TripID           string `json:"trip_id" validate:"required"`
	EndDate          string `json:"end_date" validate:"required"`
	EndKM            string `json:"end_km" validate:"required"`
	ReturnedQuantity string `json:"returned_quantity" validate:"required"`
	DamagedQuantity  string `json:"damaged_quantity" validate:"required"`
	Remarks          string `json:"remarks" validate:"required"`
	Status           string `json:"status" validate:"required"`
}

type TripResponse struct {
	TripID              string             `json:"trip_id"`
	UserID              string             `json:"user_id"`
	User                coreModel.User     `json:"user"`
	StartDate           string             `json:"start_date" `
	EndDate             string             `json:"end_date"`
	VehicleNumber       string             `json:"vehicle_number"`
	VehicleType         string             `json:"vehicle_type"`
	VehicleName         string             `json:"vehicle_name"`
	DriverName          string             `json:"driver_name"`
	DriverContactNumber string             `json:"driver_contact_number"`
	DriverProof         string             `json:"driver_proof"`
	StartKM             string             `json:"start_km"`
	EndKM               string             `json:"end_km"`
	LoadedQuantity      string             `json:"loaded_quantity"`
	ReturnedQuantity    string             `json:"returned_quantity"`
	DamagedQuantity     string             `json:"damaged_quantity"`
	Status              string             `json:"status"`
	Items               []TripItemResponse `json:"item" `
}

type TripItemResponse struct {
	TripItemID   string                     `json:"trip_item_id"`
	InvoiceID    string                     `json:"invoice_id"`
	OrderID      string                     `json:"order_id"`
	Order        orderDto.OrderResponse     `json:"order"`
	OutletID     string                     `json:"outlet_id"`
	Outlet       outletDto.OutletResponse   `json:"outlet"`
	ProductID    string                     `json:"product_id"`
	Product      productDto.ProductResponse `json:"product"`
	Quantity     string                     `json:"quantity"`
	FreeQuantity string                     `json:"free_quantity"`
	Remarks      string                     `json:"remarks"`
	Status       string                     `json:"status"`
}
