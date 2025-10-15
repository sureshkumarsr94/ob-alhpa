package distributor_service

import (
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/dto"
	distributorDto "infopack.co.in/offybox/app/dto/distributor"
	userService "infopack.co.in/offybox/app/services/user"
)

// DistributorService defines the interface for distributor-related operations
type DistributorService interface {
	SaveDistributor(tx *gorm.DB, params distributorDto.DistributorRequest) (employeeId string, handle dto.HandleError)

	GetDistributor(employeeId string, userService userService.UserService) (payload distributorDto.DistributorResponse)

	SaveDistributorUser(tx *gorm.DB, params distributorDto.DistributorUserRequest) (employeeId string, handle dto.HandleError)

	GetDistributorUsers()
}

// distributorService is an implementation of DistributorService
type distributorService struct{}

// NewDistributorService returns a new instance of DistributorService
func NewDistributorService() DistributorService {
	return &distributorService{}
}
