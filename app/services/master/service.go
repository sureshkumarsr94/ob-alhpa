package master_service

import (
	"infopack.co.in/offybox/app/dto"
	masterDto "infopack.co.in/offybox/app/dto/master"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
)

// MasterService defines the interface for master-related operations
type MasterService interface {

	// SaveRole save the role
	SaveRole(params masterDto.RoleRequest, userObject coreModel.User) (roleID string, handle dto.HandleError)

	// GetRoleObject role list
	GetRoleObject(roleID string) (payload masterDto.RoleResponse)

	// SaveWarehouse save the Warehouse
	SaveWarehouse(params masterDto.WarehouseRequest, userObject coreModel.User) (warehouse entityModel.Warehouse, handle dto.HandleError)

	// GetWarehouseObject Warehouse list
	GetWarehouseObject(warehouseID string) (payload masterDto.WarehouseResponse)
}

// masterService is an implementation of MasterService
type masterService struct{}

// NewMasterService returns a new instance of MasterService
func NewMasterService() MasterService {
	return &masterService{}
}
