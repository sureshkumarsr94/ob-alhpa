package master_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	"infopack.co.in/offybox/app/database"
	masterDto "infopack.co.in/offybox/app/dto/master"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
	masterService "infopack.co.in/offybox/app/services/master"
	userService "infopack.co.in/offybox/app/services/user"
	"net/http"
)

func RoleList(c *fiber.Ctx, userService userService.UserService, masterService masterService.MasterService) error {
	user := userService.GetUserModel(c)
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.RoleColumns, key) {
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
			Key:       coreModel.RoleColumns.DistributorID,
			Condition: "=",
			Value:     distributorUser.DistributorID,
		})
	}

	role := coreModel.Role{}
	list, _ := role.FindAll(conditions)
	var response []masterDto.RoleResponse
	for _, role = range list {
		response = append(response, masterService.GetRoleObject(role.ID))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateRole(c *fiber.Ctx, userService userService.UserService, masterService masterService.MasterService) error {
	var params masterDto.RoleRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	user := userService.GetUserModel(c)
	tx := database.MysqlDB.Begin()
	roleID, hle := masterService.SaveRole(params, user)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   masterService.GetRoleObject(roleID),
	})
}

func WarehouseList(c *fiber.Ctx, userService userService.UserService, masterService masterService.MasterService) error {
	user := userService.GetUserModel(c)
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(entityModel.WarehouseColumns, key) {
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
			Key:       entityModel.WarehouseColumns.DistributorID,
			Condition: "=",
			Value:     distributorUser.DistributorID,
		})
	}

	warehouse := entityModel.Warehouse{}
	list, _ := warehouse.FindAll(conditions)
	var response []masterDto.WarehouseResponse
	for _, warehouse = range list {
		response = append(response, masterService.GetWarehouseObject(warehouse.ID))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateWarehouse(c *fiber.Ctx, userService userService.UserService, masterService masterService.MasterService) error {
	var params masterDto.WarehouseRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	user := userService.GetUserModel(c)
	tx := database.MysqlDB.Begin()
	warehouse, hle := masterService.SaveWarehouse(params, user)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   masterService.GetWarehouseObject(warehouse.ID),
	})
}
