package territory_service

import (
	"infopack.co.in/offybox/app/dto"
	territoryDto "infopack.co.in/offybox/app/dto/territory"
	coreModel "infopack.co.in/offybox/app/models/core"
)

// TerritoryService defines the interface for territory-related operations
type TerritoryService interface {

	// GetTerritoryTypeObject retrieves the territory type object
	GetTerritoryTypeObject(territoryType coreModel.TerritoryType) (payload territoryDto.TerritoryTypeResponse)

	// SaveTerritoryType save the territory type
	SaveTerritoryType(params territoryDto.TerritoryTypeRequest, userId string) (territoryType coreModel.TerritoryType, handle dto.HandleError)

	// GetTerritoryObject retrieves the territory object
	GetTerritoryObject(territory coreModel.Territory) (payload territoryDto.TerritoryResponse)

	// SaveTerritory save the territory
	SaveTerritory(params territoryDto.TerritoryRequest, userId string) (territory coreModel.Territory, handle dto.HandleError)
}

// territoryService is an implementation of TerritoryService
type territoryService struct{}

// NewTerritoryService returns a new instance of TerritoryService
func NewTerritoryService() TerritoryService {
	return &territoryService{}
}
