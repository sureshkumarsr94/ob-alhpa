package outlet_service

import (
	"infopack.co.in/offybox/app/dto"
	outletDto "infopack.co.in/offybox/app/dto/outlet"
	coreModel "infopack.co.in/offybox/app/models/core"
	userService "infopack.co.in/offybox/app/services/user"
)

// OutletService defines the interface for outlet-related operations
type OutletService interface {
	SaveOutletCategory(params outletDto.OutletCategoryRequest, userObject coreModel.User) (outletCategoryID string, handle dto.HandleError)

	GetOutletCategory(outletCategoryID string) (payload outletDto.OutletCategoryResponse)

	SaveOutlet(params outletDto.OutletRequest, userObject coreModel.User) (outletID string, handle dto.HandleError)

	GetOutlet(outletID string, userService userService.UserService) (payload outletDto.OutletResponse)

	SaveOutletAddress(params outletDto.OutletAddressRequest, userObject coreModel.User) (outletAddressID string, handle dto.HandleError)

	GetOutletAddress(outletAddressID string) (payload outletDto.OutletAddressResponse)
}

// outletService is an implementation of OutletService
type outletService struct{}

// NewOutletService returns a new instance of OutletService
func NewOutletService() OutletService {
	return &outletService{}
}
