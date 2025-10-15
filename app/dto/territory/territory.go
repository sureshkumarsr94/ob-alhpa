package territory_dto

type TerritoryTypeRequest struct {
	TerritoryTypeId       string `json:"territory_type_id"`
	Code                  string `json:"code" validate:"required"`
	Name                  string `json:"name" validate:"required"`
	Description           string `json:"description"`
	ParentTerritoryTypeId string `json:"parent_territory_type_id"`
	Status                string `json:"status" validate:"required"`
}

type TerritoryTypeResponse struct {
	TerritoryTypeId       string                 `json:"territory_type_id"`
	Code                  string                 `json:"code"`
	Name                  string                 `json:"name"`
	Description           string                 `json:"description"`
	Level                 string                 `json:"level"`
	ParentTerritoryTypeId string                 `json:"parent_territory_type_id"`
	ParentTerritoryType   *TerritoryTypeResponse `json:"parent_territory_type,omitempty"`
	Status                string                 `json:"status"`
}

type TerritoryRequest struct {
	TerritoryId       string `json:"territory_id"`
	Code              string `json:"code" validate:"required"`
	Name              string `json:"name" validate:"required"`
	Description       string `json:"description"`
	TerritoryTypeId   string `json:"territory_type_id" validate:"required"`
	ParentTerritoryId string `json:"parent_territory_id"`
	LocationType      string `json:"location_type"`
	LocationTypeValue string `json:"location_type_value"`
	Status            string `json:"status" validate:"required"`
}

type TerritoryResponse struct {
	TerritoryId       string                `json:"territory_id"`
	Code              string                `json:"code"`
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	TerritoryTypeId   string                `json:"territory_type_id"`
	TerritoryType     TerritoryTypeResponse `json:"territory_type,omitempty"`
	ParentTerritoryId string                `json:"parent_territory_id"`
	ParentTerritory   *TerritoryResponse    `json:"parent_territory,omitempty"`
	LocationType      string                `json:"location_type"`
	LocationTypeValue string                `json:"location_type_value"`
	Status            string                `json:"status" validate:"required"`
}
