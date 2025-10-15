package distributor_service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/passwordutil"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/database"
	"infopack.co.in/offybox/app/dto"
	distributorDto "infopack.co.in/offybox/app/dto/distributor"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
	userService "infopack.co.in/offybox/app/services/user"
)

func (s *distributorService) SaveDistributor(tx *gorm.DB, params distributorDto.DistributorRequest) (distributorId string, handle dto.HandleError) {

	user := coreModel.User{}
	distributorUser := entityModel.DistributorUser{}
	// employee duplicate validation

	user, _ = user.FindByEmailOrMobile(constants.UserTypeDistributor, params.Email, params.Mobile)
	if len(user.ID) > 0 {
		checkDistributorUser, _ := distributorUser.FindByUserId(user.ID)
		if checkDistributorUser.DistributorID != params.DistributorID {
			handle = dto.HandleError{
				Status: -1,
				Errors: fmt.Errorf("distributor already exists"),
			}
			return
		}
	}

	distributor := entityModel.Distributor{}
	distributor, _ = distributor.FindByPrimaryKey(params.DistributorID)

	newDistributor := false
	if len(distributor.ID) == 0 {
		newDistributor = true
		distributor.ID = uuid.New().String()
		distributor.Code = params.Code
		if distributor.Code == "" {
			var preCode string
			recode, _ := distributor.FindByLastInsertRuleRecode()
			preCode = recode.Code
			code, _ := utility.GenerateCode(constants.MasterPseudo, constants.PrefixDistributor, &preCode, nil)
			distributor.Code = code
		}
	}

	distributor.Name = params.Name
	distributor.Email = params.Email
	distributor.Mobile = params.Mobile
	distributor.PointOfContact = params.PointOfContact
	distributor.Status = params.Status

	if len(user.ID) == 0 {
		user.ID = uuid.New().String()
		user.Type = constants.UserTypeDistributor
	}

	user.FirstName, user.LastName = utility.ExtractFirstAndLastName(params.PointOfContact)
	user.Username = params.Name
	user.Email = params.Email
	user.Mobile = params.Mobile
	user.Status = params.Status

	if err := tx.Save(&user).Error; err != nil {
		handle = dto.HandleError{
			Status: -2,
			Errors: err,
		}
		return
	}

	if err := tx.Save(&distributor).Error; err != nil {
		handle = dto.HandleError{
			Status: -1,
			Errors: err,
		}
	}

	if newDistributor {
		distributorUser := entityModel.DistributorUser{}
		role := coreModel.Role{}

		role.ID = uuid.New().String()
		role.UserType = constants.UserTypeDistributor
		role.Code = constants.UserRoleAdmin
		role.Name = constants.UserRoleAdmin
		role.DataAccess = constants.DataAccessAll
		role.DistributorID = null.StringFrom(distributor.ID)
		if err := tx.Save(&role).Error; err != nil {
			handle = dto.HandleError{
				Status: -2,
				Errors: err,
			}
			return
		}

		userRole := coreModel.UserRole{}
		userRole.RoleID = role.ID
		userRole.UserID = user.ID
		userRole.IsPrimary = coreModel.UserRoleIsPrimary
		if err := tx.Save(&userRole).Error; err != nil {
			handle = dto.HandleError{
				Status: -2,
				Errors: err,
			}
			return
		}

		distributorUser.ID = uuid.New().String()
		distributorUser.DistributorID = distributor.ID
		distributorUser.UserID = user.ID
		distributorUser.RoleID = role.ID

		if err := tx.Save(&distributorUser).Error; err != nil {
			handle = dto.HandleError{
				Status: -5,
				Errors: err,
			}
			return
		}

		user.AssociateUserId = distributorUser.ID

		if err := tx.Save(&user).Error; err != nil {
			handle = dto.HandleError{
				Status: -5,
				Errors: err,
			}
			return
		}
	}

	err := tx.Commit().Error
	if err != nil {
		handle = dto.HandleError{
			Status: -8,
			Errors: err,
		}
	}
	distributorId = distributor.ID
	return
}

func (s *distributorService) GetDistributor(distributorId string, userService userService.UserService) (payload distributorDto.DistributorResponse) {

	distributor := entityModel.Distributor{}
	distributor, _ = distributor.FindByPrimaryKey(distributorId)
	if distributor.ID != "" {
		payload.DistributorID = distributor.ID
		payload.Name = distributor.Name
		payload.Code = distributor.Code
		payload.Mobile = distributor.Mobile
		payload.Email = distributor.Email
		payload.Status = distributor.Status
		payload.PointOfContact = distributor.PointOfContact

		role := coreModel.Role{}
		role, _ = role.FindByDistributorAndRole(distributorId, constants.UserRoleAdmin)

		var condition []database.WhereCondition
		condition = append(condition, database.WhereCondition{
			Key:       entityModel.DistributorUserColumns.DistributorID,
			Condition: "=",
			Value:     distributor.ID,
		})

		condition = append(condition, database.WhereCondition{
			Key:       entityModel.DistributorUserColumns.RoleID,
			Condition: "=",
			Value:     role.ID,
		})

		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindOneByCondition(condition)

		payload.UserDetail = userService.GetUserObject(distributorUser.UserID)

	}
	return
}

func (s *distributorService) GetDistributorUsers() {}

func (s *distributorService) SaveDistributorUser(tx *gorm.DB, params distributorDto.DistributorUserRequest) (userId string, handle dto.HandleError) {

	user := coreModel.User{}
	role := coreModel.Role{}
	distributorUser := entityModel.DistributorUser{}

	role, _ = role.FindByPrimaryKey(params.RoleID)
	if len(role.ID) == 0 {
		handle = dto.HandleError{
			Status: -1,
			Errors: fmt.Errorf("role details not found"),
		}
		return
	}
	// employee duplicate validation
	user, _ = user.FindByEmailOrMobile(constants.UserTypeDistributor, params.Email, params.Mobile)

	if len(user.ID) == 0 {
		user.ID = uuid.New().String()
		user.Type = constants.UserTypeDistributor
	}

	user.Username = params.FirstName + " " + params.LastName
	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Email = params.Email
	user.Mobile = params.Mobile
	user.Status = params.Status
	if len(params.Password) > 0 {
		user.Password = passwordutil.HashPassword(params.Password)
	}

	if err := tx.Save(&user).Error; err != nil {
		handle = dto.HandleError{
			Status: -2,
			Errors: err,
		}
		return
	}

	userRole := coreModel.UserRole{}
	userRole, _ = userRole.FindByUser(userId)
	userRole.RoleID = role.ID
	userRole.UserID = user.ID
	userRole.IsPrimary = coreModel.UserRoleIsPrimary
	if err := tx.Save(&userRole).Error; err != nil {
		handle = dto.HandleError{
			Status: -3,
			Errors: err,
		}
		return
	}

	if len(user.ID) > 0 {
		checkDistributorUser, _ := distributorUser.FindByUserId(user.ID)
		if checkDistributorUser.DistributorID != params.DistributorID {
			distributorUser.ID = uuid.New().String()
			distributorUser.DistributorID = params.DistributorID
			distributorUser.UserID = user.ID
			distributorUser.RoleID = params.RoleID
			if err := tx.Save(&distributorUser).Error; err != nil {
				handle = dto.HandleError{
					Status: -4,
					Errors: err,
				}
			}
			user.AssociateUserId = distributorUser.ID

			if err := tx.Save(&user).Error; err != nil {
				handle = dto.HandleError{
					Status: -5,
					Errors: err,
				}
				return
			}
		}
	}

	err := tx.Commit().Error
	if err != nil {
		handle = dto.HandleError{
			Status: -4,
			Errors: err,
		}
	}
	userId = user.ID
	return
}
