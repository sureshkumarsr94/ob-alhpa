package master_dto

type WarehouseRequest struct {
	WarehouseID string `json:"warehouse_id"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required"`
}

type WarehouseResponse struct {
	WarehouseID string `json:"warehouse_id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
