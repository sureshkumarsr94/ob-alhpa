package stock_service

import (
	"infopack.co.in/offybox/app/dto"
	stockDto "infopack.co.in/offybox/app/dto/stock"
)

// StockService defines the interface for journey plan-related operations
type StockService interface {

	// GetStock retrieves the  stock object
	GetJStock(stockID string) (payload stockDto.StockResponse)

	// SaveStock save the journey plan
	SaveStock(params stockDto.StockRequest) (stockID string, handle dto.HandleError)
}

// stockService is an implementation of StockService
type stockService struct{}

// NewStockService returns a new instance of StockService
func NewStockService() StockService {
	return &stockService{}
}
