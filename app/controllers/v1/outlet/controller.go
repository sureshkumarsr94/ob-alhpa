package outlet_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	"infopack.co.in/offybox/app/database"
	outletDto "infopack.co.in/offybox/app/dto/outlet"
	entityModel "infopack.co.in/offybox/app/models/entity"
	outletService "infopack.co.in/offybox/app/services/outlet"
	userService "infopack.co.in/offybox/app/services/user"
	"net/http"
)

func OutletList(c *fiber.Ctx, userService userService.UserService, outletService outletService.OutletService) error {
	user := userService.GetUserModel(c)
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(entityModel.OutletColumns, key) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -1,
				"error":  fmt.Sprintf("Invalid column %s", key),
			})
		}
		condition := database.WhereCondition{}
		condition.Key = key
		condition.Condition = "="
		condition.Value = value
		conditions = append(conditions, condition)
	}

	// Add condition to match the user type as 'Distributor'
	if user.Type == constants.UserTypeDistributor {
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(user.ID)
		if len(distributorUser.ID) == 0 {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -2,
				"error":  "Distributor Information not Found",
			})
		}
		conditions = append(conditions, database.WhereCondition{
			Key:       entityModel.OutletColumns.DistributorID,
			Condition: "=",
			Value:     distributorUser.DistributorID,
		})
	} else {
		conditions = append(conditions, database.WhereCondition{
			Key:       entityModel.OutletColumns.DistributorID,
			Condition: "IS",
			Value:     "NULL",
		})
	}

	outlet := entityModel.Outlet{}
	list, _ := outlet.FindAll(conditions)
	var response []outletDto.OutletResponse
	for _, outlet = range list {
		response = append(response, outletService.GetOutlet(outlet.ID, userService))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateOutlet(c *fiber.Ctx, userService userService.UserService, outletService outletService.OutletService) error {
	var params outletDto.OutletRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	user := userService.GetUserModel(c)
	tx := database.MysqlDB.Begin()
	outletID, hle := outletService.SaveOutlet(params, user)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   outletService.GetOutlet(outletID, userService),
	})
}

func OutletCategoryList(c *fiber.Ctx, userService userService.UserService, outletService outletService.OutletService) error {
	user := userService.GetUserModel(c)
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(entityModel.OutletCategoryColumns, key) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -1,
				"error":  fmt.Sprintf("Invalid column %s", key),
			})
		}
		condition := database.WhereCondition{}
		condition.Key = key
		condition.Condition = "="
		condition.Value = value
		conditions = append(conditions, condition)
	}

	// Add condition to match the user type as 'Distributor'
	if user.Type == constants.UserTypeDistributor {
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(user.ID)
		if len(distributorUser.ID) == 0 {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -2,
				"error":  "Distributor Information not Found",
			})
		}
		conditions = append(conditions, database.WhereCondition{
			Key:       entityModel.OutletCategoryColumns.DistributorID,
			Condition: "=",
			Value:     distributorUser.DistributorID,
		})
	} else {
		conditions = append(conditions, database.WhereCondition{
			Key:       entityModel.OutletCategoryColumns.DistributorID,
			Condition: "IS",
			Value:     "NULL",
		})
	}

	outletCategory := entityModel.OutletCategory{}
	list, _ := outletCategory.FindAll(conditions)
	var response []outletDto.OutletCategoryResponse
	for _, outletCategory = range list {
		response = append(response, outletService.GetOutletCategory(outletCategory.ID))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateOutletCategory(c *fiber.Ctx, userService userService.UserService, outletService outletService.OutletService) error {
	var params outletDto.OutletCategoryRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	user := userService.GetUserModel(c)
	tx := database.MysqlDB.Begin()
	outletCategoryID, hle := outletService.SaveOutletCategory(params, user)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   outletService.GetOutletCategory(outletCategoryID),
	})
}

func OutletAddressList(c *fiber.Ctx, outletService outletService.OutletService) error {
	outletID := c.Params("outletID")
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(entityModel.OutletAddressColumns, key) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -1,
				"error":  fmt.Sprintf("Invalid column %s", key),
			})
		}
		condition := database.WhereCondition{}
		condition.Key = key
		condition.Condition = "="
		condition.Value = value
		conditions = append(conditions, condition)
	}

	conditions = append(conditions, database.WhereCondition{
		Key:       entityModel.OutletAddressColumns.OutletID,
		Condition: "=",
		Value:     outletID,
	})

	outletAddress := entityModel.OutletAddress{}
	list, _ := outletAddress.FindAll(conditions)
	var response []outletDto.OutletAddressResponse
	for _, outletAddress = range list {
		response = append(response, outletService.GetOutletAddress(outletAddress.ID))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateOutletAddress(c *fiber.Ctx, userService userService.UserService, outletService outletService.OutletService) error {
	var params outletDto.OutletAddressRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	user := userService.GetUserModel(c)
	tx := database.MysqlDB.Begin()
	outletAddressID, hle := outletService.SaveOutletAddress(params, user)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   outletService.GetOutletAddress(outletAddressID),
	})
}
