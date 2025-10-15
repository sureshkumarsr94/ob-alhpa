package employee_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	"infopack.co.in/offybox/app/database"
	employeeDto "infopack.co.in/offybox/app/dto/employee"
	coreModel "infopack.co.in/offybox/app/models/core"
	employeeService "infopack.co.in/offybox/app/services/employee"
	"net/http"
)

func EmployeeList(c *fiber.Ctx, employeeService employeeService.EmployeeService) error {
	queries := c.Queries()
	var conditions []database.WhereCondition
	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.EmployeeColumns, key) {
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
	employee := coreModel.Employee{}
	list, _ := employee.FindAll(conditions)
	var response []employeeDto.EmployeeResponse
	for _, employee = range list {
		response = append(response, employeeService.GetEmployee(employee.ID))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})
}

func CreateUpdateEmployee(c *fiber.Ctx, employeeService employeeService.EmployeeService) error {
	var params employeeDto.EmployeeRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}

	tx := database.MysqlDB.Begin()
	employeeID, hle := employeeService.SaveEmployee(tx, params)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   employeeService.GetEmployee(employeeID),
	})
}
