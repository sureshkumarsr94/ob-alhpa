package territory_service

import (
	"fmt"
	"github.com/guregu/null"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/dto"
	territoryDto "infopack.co.in/offybox/app/dto/territory"
	coreModel "infopack.co.in/offybox/app/models/core"
)

func (s *territoryService) GetTerritoryTypeObject(territoryType coreModel.TerritoryType) (payload territoryDto.TerritoryTypeResponse) {

	if territoryType.ParentID.String != "" {
		parentTerritoryType, _ := territoryType.FindByPrimaryKey(territoryType.ParentID.String)
		parentPayload := territoryDto.TerritoryTypeResponse{
			TerritoryTypeId: parentTerritoryType.ID,
			Code:            parentTerritoryType.Code,
			Level:           utility.ToString(parentTerritoryType.Level),
			Name:            parentTerritoryType.Name,
			Description:     parentTerritoryType.Description,
			Status:          parentTerritoryType.Status,
		}
		payload.ParentTerritoryType = &parentPayload

	}
	payload.TerritoryTypeId = territoryType.ID
	payload.ParentTerritoryTypeId = territoryType.ParentID.String
	payload.Code = territoryType.Code
	payload.Name = territoryType.Name
	payload.Level = utility.ToString(territoryType.Level)
	payload.Description = territoryType.Description
	payload.Status = territoryType.Status

	return
}

func (s *territoryService) SaveTerritoryType(params territoryDto.TerritoryTypeRequest, userId string) (territoryType coreModel.TerritoryType, handle dto.HandleError) {

	if params.TerritoryTypeId != "" {
		territoryType, _ = territoryType.FindByPrimaryKey(params.TerritoryTypeId)
		if territoryType.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid territory type")
			return
		}
	}

	if params.ParentTerritoryTypeId != "" {
		parentTerritoryType, _ := territoryType.FindByPrimaryKey(params.ParentTerritoryTypeId)
		if parentTerritoryType.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid parent territory type")
			return
		}
		territoryType.ParentID = null.StringFrom(params.ParentTerritoryTypeId)
		territoryType.Level = parentTerritoryType.Level + 1
	} else {
		territoryType.Level = 1
	}

	territoryType.Code = params.Code
	territoryType.Name = params.Name
	territoryType.Description = params.Description
	territoryType.Status = params.Status
	territoryType, err := territoryType.Save()
	if err != nil {
		handle.Status = -1
		handle.Errors = err
		return
	}

	return
}

func (s *territoryService) GetTerritoryObject(territory coreModel.Territory) (payload territoryDto.TerritoryResponse) {

	if territory.ParentID.String != "" {
		parentTerritory, _ := territory.FindByPrimaryKey(territory.ParentID.String)
		parentPayload := territoryDto.TerritoryResponse{
			TerritoryId:       parentTerritory.ID,
			Code:              parentTerritory.Code,
			Name:              parentTerritory.Name,
			Description:       parentTerritory.Description,
			Status:            parentTerritory.Status,
			ParentTerritoryId: parentTerritory.ParentID.String,
			TerritoryTypeId:   parentTerritory.TerritoryTypeID,
		}
		payload.ParentTerritory = &parentPayload
	}
	payload.TerritoryId = territory.ID
	payload.Code = territory.Code
	payload.Name = territory.Name
	payload.Description = territory.Description
	payload.Status = territory.Status
	payload.ParentTerritoryId = territory.ParentID.String
	payload.TerritoryTypeId = territory.TerritoryTypeID
	payload.TerritoryType.TerritoryTypeId = territory.TerritoryType.ID
	payload.TerritoryType.Code = territory.TerritoryType.Code
	payload.TerritoryType.Name = territory.TerritoryType.Name
	payload.TerritoryType.Description = territory.TerritoryType.Description
	payload.TerritoryType.Status = territory.TerritoryType.Status
	payload.TerritoryType.ParentTerritoryTypeId = territory.TerritoryType.ParentID.String
	return
}

func (s *territoryService) SaveTerritory(params territoryDto.TerritoryRequest, userId string) (territory coreModel.Territory, handle dto.HandleError) {

	if params.TerritoryId != "" {
		territory, _ = territory.FindByPrimaryKey(params.TerritoryId)
		if territory.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid territory")
			return
		}
	}

	if params.ParentTerritoryId != "" {
		parentTerritory, _ := territory.FindByPrimaryKey(params.ParentTerritoryId)
		if parentTerritory.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid parent territory")
			return
		}
		territory.ParentID = null.StringFrom(params.ParentTerritoryId)
	}

	if params.TerritoryTypeId != "" {
		territoryType := coreModel.TerritoryType{}
		territoryType, _ = territoryType.FindByPrimaryKey(params.TerritoryTypeId)
		if territoryType.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid territory type")
			return
		}
		territory.TerritoryTypeID = territoryType.ID
	}

	territory.Code = params.Code
	territory.Name = params.Name
	territory.Description = params.Description
	territory.Status = params.Status
	territory, err := territory.Save()
	if err != nil {
		handle.Status = -1
		handle.Errors = err
		return
	}

	return
}
