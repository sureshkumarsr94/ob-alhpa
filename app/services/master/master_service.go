package master_service

import (
	"fmt"
	"github.com/guregu/null"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/dto"
	masterDto "infopack.co.in/offybox/app/dto/master"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
)

func (*masterService) SaveRole(params masterDto.RoleRequest, userObject coreModel.User) (roleID string, handle dto.HandleError) {
	role := coreModel.Role{}
	role, _ = role.FindByPrimaryKey(params.RoleID)
	role.Code = params.Code
	role.Name = params.Name
	role.Description = params.Description
	role.Status = params.Status
	role.DataAccess = params.DataAccess
	if len(params.ParentRoleID) > 0 {
		parentRole, _ := role.FindByPrimaryKey(params.ParentRoleID)
		if len(parentRole.ID) == 0 {
			handle = dto.HandleError{
				Status: -1,
				Errors: fmt.Errorf("parent role not found"),
			}
			return
		}
	}

	if userObject.Type == constants.UserTypeDistributor {
		role.UserType = constants.UserTypeDistributor
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(userObject.ID)
		if len(distributorUser.ID) == 0 {
			handle = dto.HandleError{
				Status: -3,
				Errors: fmt.Errorf("distributor user mapping not found"),
			}
			return
		}
		role.DistributorID = null.StringFrom(distributorUser.DistributorID)
	} else {
		role.UserType = constants.UserTypeEmployee
	}
	role, err := role.Save()
	if err != nil {
		handle = dto.HandleError{
			Status: -4,
			Errors: err,
		}
		return
	}
	roleID = role.ID
	return
}

func (*masterService) GetRoleObject(roleID string) (payload masterDto.RoleResponse) {
	role := coreModel.Role{}
	role, _ = role.FindByPrimaryKey(roleID)
	payload.RoleID = role.ID
	payload.Code = role.Code
	payload.Name = role.Name
	payload.Description = role.Description
	payload.DataAccess = role.DataAccess
	payload.Status = role.Status
	payload.ParentRoleID = role.RoleID.Ptr()
	if !role.RoleID.IsZero() {
		role, _ = role.FindByPrimaryKey(role.RoleID.String)
		parentRolePayload := masterDto.RoleResponse{
			RoleID:       role.RoleID.String,
			Code:         role.Code,
			Name:         role.Name,
			DataAccess:   role.DataAccess,
			Description:  role.Description,
			Status:       role.Status,
			ParentRoleID: role.RoleID.Ptr(),
		}
		payload.ParentRole = &parentRolePayload
	}
	return
}

func (*masterService) SaveWarehouse(params masterDto.WarehouseRequest, userObject coreModel.User) (warehouse entityModel.Warehouse, handle dto.HandleError) {

	warehouse, _ = warehouse.FindByPrimaryKey(params.WarehouseID)
	warehouse.Code = params.Code
	warehouse.Name = params.Name
	warehouse.Description = params.Description
	warehouse.Status = params.Status

	if userObject.Type == constants.UserTypeDistributor {
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(userObject.ID)
		if len(distributorUser.ID) == 0 {
			handle = dto.HandleError{
				Status: -3,
				Errors: fmt.Errorf("distributor user mapping not found"),
			}
			return
		}
		warehouse.DistributorID = null.StringFrom(distributorUser.DistributorID)
	}
	warehouse, err := warehouse.Save()

	if err != nil {
		handle = dto.HandleError{
			Status: -4,
			Errors: err,
		}
		return
	}

	return
}

func (*masterService) GetWarehouseObject(warehouseID string) (payload masterDto.WarehouseResponse) {
	warehouse := entityModel.Warehouse{}
	warehouse, _ = warehouse.FindByPrimaryKey(warehouseID)
	payload.WarehouseID = warehouse.ID
	payload.Code = warehouse.Code
	payload.Name = warehouse.Name
	payload.Status = warehouse.Status
	payload.Description = warehouse.Description
	return
}
