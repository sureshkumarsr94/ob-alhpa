package distributor_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	"infopack.co.in/offybox/app/database"
	distributorDto "infopack.co.in/offybox/app/dto/distributor"
	userDto "infopack.co.in/offybox/app/dto/user"
	entityModel "infopack.co.in/offybox/app/models/entity"
	distributorService "infopack.co.in/offybox/app/services/distributor"
	userService "infopack.co.in/offybox/app/services/user"
	"net/http"
)

func GetDistributors(c *fiber.Ctx, distributorService distributorService.DistributorService, userService userService.UserService) error {
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(entityModel.DistributorColumns, key) {
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

	distributor := entityModel.Distributor{}
	list, _ := distributor.FindAll(conditions)
	var response []distributorDto.DistributorResponse
	for _, d := range list {
		response = append(response, distributorService.GetDistributor(d.ID, userService))
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateDistributor(c *fiber.Ctx, distributorService distributorService.DistributorService, userService userService.UserService) error {
	var params distributorDto.DistributorRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}

	tx := database.MysqlDB.Begin()
	distributorId, hle := distributorService.SaveDistributor(tx, params)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   distributorService.GetDistributor(distributorId, userService),
	})
}

func GetDistributor(c *fiber.Ctx, distributorService distributorService.DistributorService, userService userService.UserService) error {
	data := distributorService.GetDistributor(c.Params("distributorId"), userService)
	if data.DistributorID == "" {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  "distributor details not found",
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   data,
	})
}

func GetDistributorUsers(c *fiber.Ctx, userService userService.UserService) error {
	user := userService.GetUserModel(c)
	queries := c.Queries()
	var conditions []database.WhereCondition

	for key, value := range queries {
		if !utility.IsValidColumn(entityModel.DistributorUserColumns, key) {
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
			Key:       entityModel.DistributorUserColumns.DistributorID,
			Condition: "=",
			Value:     distributorUser.DistributorID,
		})
	}

	distributorUser := entityModel.DistributorUser{}
	list, _ := distributorUser.FindAll(conditions)
	var response []userDto.UserObject
	for _, d := range list {
		response = append(response, userService.GetUserObject(d.UserID))
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateDistributorUser(c *fiber.Ctx, distributorService distributorService.DistributorService, userService userService.UserService) error {
	var params distributorDto.DistributorUserRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}

	tx := database.MysqlDB.Begin()
	userID, hle := distributorService.SaveDistributorUser(tx, params)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   userService.GetUserObject(userID),
	})
}
