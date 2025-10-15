package territory_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	db "infopack.co.in/offybox/app/database"
	territoryDto "infopack.co.in/offybox/app/dto/territory"
	coreModel "infopack.co.in/offybox/app/models/core"
	territoryService "infopack.co.in/offybox/app/services/territory"
	userService "infopack.co.in/offybox/app/services/user"
	"net/http"
)

func TerritoryTypeList(c *fiber.Ctx, userService userService.UserService, territoryService territoryService.TerritoryService) error {
	//user := userService.GetUserObject(c)
	territoryType := coreModel.TerritoryType{}
	queries := c.Queries()
	var conditions []db.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.CountryColumns, key) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -1,
				"error":  fmt.Sprintf("Invalid column %s", key),
			})
		}
		condition := db.WhereCondition{}
		condition.Key = key
		condition.Condition = "="
		condition.Value = value
		conditions = append(conditions, condition)
	}

	list, err := territoryType.FindAll(conditions)
	var response []territoryDto.TerritoryTypeResponse
	for _, element := range list {
		response = append(response, territoryService.GetTerritoryTypeObject(element))
	}
	if err != nil {
		// Return a 422 Unprocessable Entity status with the validation error
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})

}

func CreateUpdateTerritoryType(c *fiber.Ctx, userService userService.UserService, territoryService territoryService.TerritoryService) error {
	user := userService.GetUserModel(c)
	var params territoryDto.TerritoryTypeRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	territoryType, hle := territoryService.SaveTerritoryType(params, user.ID)
	if hle.Status < 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  hle.Errors,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   territoryService.GetTerritoryTypeObject(territoryType),
	})
}

func TerritoryList(c *fiber.Ctx, userService userService.UserService, territoryService territoryService.TerritoryService) error {
	//user := userService.GetUserObject(c)
	territory := coreModel.Territory{}
	var andCondition []db.WhereCondition

	// Add condition to match the user type as 'Customer'
	andCondition = append(andCondition, db.WhereCondition{
		Key:       coreModel.TerritoryColumns.Status,
		Condition: "=",
		Value:     constants.StatusActive,
	})
	list, err := territory.FindAll(andCondition)
	var response []territoryDto.TerritoryResponse
	for _, element := range list {
		response = append(response, territoryService.GetTerritoryObject(element))
	}
	if err != nil {
		// Return a 422 Unprocessable Entity status with the validation error
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})

}

func CreateUpdateTerritory(c *fiber.Ctx, userService userService.UserService, territoryService territoryService.TerritoryService) error {
	user := userService.GetUserModel(c)
	var params territoryDto.TerritoryRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	territory, hle := territoryService.SaveTerritory(params, user.ID)
	if hle.Status < 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  hle.Errors,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   territoryService.GetTerritoryObject(territory),
	})
}
