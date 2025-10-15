package journey_plan_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	"infopack.co.in/offybox/app/database"
	journeyPlanDto "infopack.co.in/offybox/app/dto/journey_plan"
	saleModel "infopack.co.in/offybox/app/models/sale"
	journeyPlanService "infopack.co.in/offybox/app/services/journey_plan"
	"net/http"
)

func JourneyPlanList(c *fiber.Ctx, journeyPlanService journeyPlanService.JourneyPlanService) error {
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(saleModel.JourneyPlanColumns, key) {
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

	journeyPlan := saleModel.JourneyPlan{}
	list, _ := journeyPlan.FindAll(conditions)
	var response []journeyPlanDto.JourneyPlanResponse
	for _, journeyPlan = range list {
		response = append(response, journeyPlanService.GetJourneyPlan(journeyPlan.ID))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateJourneyPlan(c *fiber.Ctx, journeyPlanService journeyPlanService.JourneyPlanService) error {
	var params journeyPlanDto.JourneyPlanRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}

	tx := database.MysqlDB.Begin()
	journeyPlanID, hle := journeyPlanService.SaveJourneyPlan(params)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   journeyPlanService.GetJourneyPlan(journeyPlanID),
	})
}
