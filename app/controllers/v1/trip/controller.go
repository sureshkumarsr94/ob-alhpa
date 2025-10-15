package trip_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	"infopack.co.in/offybox/app/database"
	tripDto "infopack.co.in/offybox/app/dto/trip"
	entityModel "infopack.co.in/offybox/app/models/entity"
	saleModel "infopack.co.in/offybox/app/models/sale"
	tripService "infopack.co.in/offybox/app/services/trip"
	userService "infopack.co.in/offybox/app/services/user"
	"net/http"
)

func TripList(c *fiber.Ctx, userService userService.UserService, tripService tripService.TripService) error {
	user := userService.GetUserModel(c)
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(saleModel.TripColumns, key) {
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
			Key:       saleModel.TripColumns.DistributorID,
			Condition: "=",
			Value:     distributorUser.DistributorID,
		})
	} else {
		conditions = append(conditions, database.WhereCondition{
			Key:       saleModel.TripColumns.DistributorID,
			Condition: "IS",
			Value:     "NULL",
		})
	}

	trip := saleModel.Trip{}
	list, _ := trip.FindAll(conditions)
	var response []tripDto.TripResponse
	for _, trip = range list {
		response = append(response, tripService.GetTrip(trip.ID))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateTrip(c *fiber.Ctx, userService userService.UserService, tripService tripService.TripService) error {
	var params tripDto.TripRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	user := userService.GetUserModel(c)
	tx := database.MysqlDB.Begin()
	tripID, hle := tripService.SaveTrip(tx, params, user)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   tripService.GetTrip(tripID),
	})
}
