package product_dto

type ProductRequest struct {
	ProductID   string  `json:"product_id"`
	Code        string  `json:"code" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	MRP         float64 `json:"mrp" validate:"required"`
	UOM         string  `json:"uom" validate:"required"`
	Variant     string  `json:"variant" validate:"required"`
	Description string  `json:"description"`
	Barcode     string  `json:"barcode" validate:"required"`
	Status      string  `json:"status" validate:"required"`
}

type ProductResponse struct {
	ProductID   string  `json:"product_id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	MRP         float64 `json:"mrp"`
	UOM         string  `json:"uom"`
	Variant     string  `json:"variant"`
	Description string  `json:"description"`
	Barcode     string  `json:"barcode"`
	Status      string  `json:"status"`
}
