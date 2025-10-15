package location_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	"infopack.co.in/offybox/app/database"
	locationDto "infopack.co.in/offybox/app/dto/location"
	coreModel "infopack.co.in/offybox/app/models/core"
	locationService "infopack.co.in/offybox/app/services/location"
	"net/http"
)

func CountryList(c *fiber.Ctx) error {

	var country coreModel.Country
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.CountryColumns, key) {
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
	countryList, err := country.FindAll(conditions)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   countryList,
	})

}

func CreateUpdateCountry(c *fiber.Ctx, locationService locationService.LocationService) error {

	var params locationDto.CountryRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	country, hle := locationService.SaveCountry(params)
	if hle.Status < 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  hle.Errors,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   country,
	})
}

func StateList(c *fiber.Ctx) error {

	var state coreModel.State
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.StateColumns, key) {
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
	stateList, err := state.FindAll(conditions)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   stateList,
	})

}

func CreateUpdateState(c *fiber.Ctx, locationService locationService.LocationService) error {

	var params locationDto.StateRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	state, hle := locationService.SaveState(params)
	if hle.Status < 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  hle.Errors,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   state,
	})
}

func CityList(c *fiber.Ctx) error {

	var city coreModel.City
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.CityColumns, key) {
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
	cityList, err := city.FindAll(conditions)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   cityList,
	})

}

func CreateUpdateCity(c *fiber.Ctx, locationService locationService.LocationService) error {

	var params locationDto.CityRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	city, hle := locationService.SaveCity(params)
	if hle.Status < 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  hle.Errors,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   city,
	})
}

func AreaList(c *fiber.Ctx) error {
	var area coreModel.Area
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.AreaColumns, key) {
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
	areaList, err := area.FindAll(conditions)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   areaList,
	})

}

func CreateUpdateArea(c *fiber.Ctx, locationService locationService.LocationService) error {

	var params locationDto.AreaRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	area, hle := locationService.SaveArea(params)
	if hle.Status < 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  hle.Errors,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   area,
	})
}

func PinCodeSuggest(c *fiber.Ctx) error {
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.AreaColumns, key) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -1,
				"error":  fmt.Sprintf("Invalid column %s", key),
			})
		}
		condition := database.WhereCondition{}
		condition.Key = key
		if key == coreModel.AreaColumns.Pincode {
			condition.Condition = "LIKE"
			condition.Value = value + "%"
		} else {
			condition.Condition = "="
			condition.Value = value
		}

		conditions = append(conditions, condition)
	}

	pinCodeList := coreModel.Area{}
	pinCodeLists, err := pinCodeList.FindAllWithGroup(conditions)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -4,
			"error":  err,
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   pinCodeLists,
	})
}
