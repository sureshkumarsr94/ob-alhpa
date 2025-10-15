package outlet_service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/guregu/null"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/dto"
	outletDto "infopack.co.in/offybox/app/dto/outlet"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
	userService "infopack.co.in/offybox/app/services/user"
	"time"
)

func (s *outletService) SaveOutletCategory(params outletDto.OutletCategoryRequest, userObject coreModel.User) (outletCategoryID string, handle dto.HandleError) {
	outletCategory := entityModel.OutletCategory{}
	outletCategory, _ = outletCategory.FindByPrimaryKey(params.OutletCategoryID)
	outletCategory.Name = params.Name
	outletCategory.Description = params.Description
	outletCategory.Status = params.Status
	if userObject.Type == constants.UserTypeDistributor {
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(userObject.ID)
		if len(distributorUser.ID) == 0 {
			handle = dto.HandleError{
				Status: -1,
				Errors: fmt.Errorf("distributor user mapping not found"),
			}
			return
		}
		outletCategory.DistributorID = null.StringFrom(distributorUser.DistributorID)
	}
	outletCategory, err := outletCategory.Save()
	if err != nil {
		handle = dto.HandleError{
			Status: -2,
			Errors: err,
		}
		return
	}
	outletCategoryID = outletCategory.ID
	return
}

func (s *outletService) GetOutletCategory(outletCategoryID string) (payload outletDto.OutletCategoryResponse) {
	outletCategory := entityModel.OutletCategory{}
	outletCategory, _ = outletCategory.FindByPrimaryKey(outletCategoryID)
	payload.OutletCategoryID = outletCategory.ID
	payload.Name = outletCategory.Name
	payload.Description = outletCategory.Description
	payload.Status = outletCategory.Status
	return
}

func (s *outletService) SaveOutlet(params outletDto.OutletRequest, userObject coreModel.User) (outletID string, handle dto.HandleError) {

	outletCategory := entityModel.OutletCategory{}
	outletCategory, _ = outletCategory.FindByPrimaryKey(params.OutletCategoryID)
	if outletCategory.ID == "" {
		handle = dto.HandleError{
			Status: -1,
			Errors: fmt.Errorf("outlet category not found"),
		}
		return
	}

	outlet := entityModel.Outlet{}
	outlet, _ = outlet.FindByPrimaryKey(params.OutletID)
	if outlet.ID == "" {
		outlet.Type = entityModel.LeadStatus
	}
	outlet.Name = params.Name
	outlet.Status = params.Status
	outlet.Code = params.Code
	outlet.Mobile = params.Mobile
	outlet.Email = params.Email
	outlet.CreditLimit = params.CreditLimit
	outlet.IncorporationDate = time.Now()
	outlet.OutletCategoryID = outletCategory.ID
	outlet.CreatedBy = userObject.ID
	if userObject.Type == constants.UserTypeDistributor {
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(userObject.ID)
		if len(distributorUser.ID) == 0 {
			handle = dto.HandleError{
				Status: -2,
				Errors: fmt.Errorf("distributor user mapping not found"),
			}
			return
		}
		outlet.DistributorID = null.StringFrom(distributorUser.DistributorID)
	}
	outlet, err := outlet.Save()
	if err != nil {
		handle = dto.HandleError{
			Status: -3,
			Errors: err,
		}
		return
	}
	outletID = outlet.ID
	return
}

func (s *outletService) GetOutlet(outletID string, userService userService.UserService) (payload outletDto.OutletResponse) {
	outlet := entityModel.Outlet{}
	outlet, _ = outlet.FindByPrimaryKey(outletID)
	payload.OutletID = outlet.ID
	payload.Code = outlet.Code
	payload.Name = outlet.Name
	payload.Type = outlet.Type
	payload.Status = outlet.Status
	payload.Email = outlet.Email
	payload.Mobile = outlet.Mobile
	payload.IncorporationDate = outlet.IncorporationDate.Format("2006-01-02")
	payload.CreditLimit = outlet.CreditLimit
	payload.Outstanding = outlet.Outstanding
	payload.OutletCategoryID = outlet.OutletCategoryID
	payload.OutletCategory = s.GetOutletCategory(outlet.OutletCategoryID)
	payload.CreatedByUserID = outlet.CreatedBy
	payload.CreatedByUser = userService.GetUserObject(outlet.CreatedBy)
	return
}

func (s *outletService) SaveOutletAddress(params outletDto.OutletAddressRequest, userObject coreModel.User) (outletAddressID string, handle dto.HandleError) {

	// validate outlet have access to the user
	if userObject.Type == constants.UserTypeDistributor {
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(userObject.ID)
		if len(distributorUser.ID) == 0 {
			handle = dto.HandleError{
				Status: -1,
				Errors: fmt.Errorf("distributor user mapping not found"),
			}
			return
		}
		outlet := entityModel.Outlet{}
		outlet, _ = outlet.FindByPrimaryKey(params.OutletID)
		if outlet.DistributorID.String != distributorUser.DistributorID {
			handle = dto.HandleError{
				Status: -2,
				Errors: fmt.Errorf("access denied"),
			}
			return
		}
	}
	area := coreModel.Area{}
	area, _ = area.FindByPrimaryKey(params.AreaId)
	if area.ID == "" {
		handle = dto.HandleError{
			Status: -3,
			Errors: fmt.Errorf("area not found"),
		}
		return
	}
	outletAddress := entityModel.OutletAddress{}
	outletAddress, _ = outletAddress.FindByPrimaryKey(params.OutletAddressID)
	if outletAddress.ID == "" {
		outletAddress.ID = uuid.New().String()
	}
	outletAddress.AddressType = params.AddressType
	outletAddress.Address = params.Address
	outletAddress.Landmark = params.Landmark
	outletAddress.Pincode = utility.StringToInt(params.Pincode)
	outletAddress.AreaID = area.ID
	outletAddress.CityID = area.CityID
	outletAddress.StateID = area.StateID
	outletAddress.CountryID = area.CountryID
	outletAddress.OutletID = params.OutletID
	outletAddress.Status = params.Status
	outletAddress, err := outletAddress.Save()
	if err != nil {
		handle = dto.HandleError{
			Status: -3,
			Errors: err,
		}
		return
	}
	outletAddressID = outletAddress.ID
	return
}

func (s *outletService) GetOutletAddress(outletAddressID string) (payload outletDto.OutletAddressResponse) {
	outletAddress := entityModel.OutletAddress{}
	outletAddress, _ = outletAddress.FindByPrimaryKey(outletAddressID)
	payload.OutletAddressID = outletAddress.ID
	payload.OutletID = outletAddress.OutletID
	payload.Status = outletAddress.Status
	payload.AddressType = outletAddress.AddressType
	payload.Address = outletAddress.Address
	payload.Landmark = outletAddress.Landmark
	payload.Pincode = utility.ToString(outletAddress.Pincode)
	payload.AreaId = outletAddress.AreaID
	payload.CityId = outletAddress.CityID
	payload.StateId = outletAddress.StateID
	payload.CountryId = outletAddress.CountryID
	area := coreModel.Area{}
	payload.Area, _ = area.FindByPrimaryKey(outletAddress.AreaID)
	city := coreModel.City{}
	payload.City, _ = city.FindByPrimaryKey(outletAddress.CityID)
	state := coreModel.State{}
	payload.State, _ = state.FindByPrimaryKey(outletAddress.StateID)
	country := coreModel.Country{}
	payload.Country, _ = country.FindByPrimaryKey(outletAddress.CountryID)
	return
}
