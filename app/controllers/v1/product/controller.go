package product_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	db "infopack.co.in/offybox/app/database"
	productDto "infopack.co.in/offybox/app/dto/product"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
	productService "infopack.co.in/offybox/app/services/product"
	userService "infopack.co.in/offybox/app/services/user"
	"net/http"
)

func CategoryList(c *fiber.Ctx, userService userService.UserService, productService productService.ProductService) error {
	user := userService.GetUserModel(c)
	category := coreModel.Category{}
	queries := c.Queries()
	var andCondition []db.WhereCondition

	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.CategoryColumns, key) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -1,
				"error":  fmt.Sprintf("Invalid column %s", key),
			})
		}
		condition := db.WhereCondition{}
		condition.Key = key
		condition.Condition = "="
		condition.Value = value
		andCondition = append(andCondition, condition)
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
		andCondition = append(andCondition, db.WhereCondition{
			Key:       coreModel.CategoryColumns.DistributorID,
			Condition: "=",
			Value:     distributorUser.DistributorID,
		})
	}

	list, err := category.FindAll(andCondition)
	if err != nil {
		// Return a 422 Unprocessable Entity status with the validation error
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -3,
			"error":  err,
		})
	}

	var response []productDto.CategoryResponse
	for _, element := range list {
		response = append(response, productService.GetCategoryObject(element.ID))
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   response,
	})

}

func CreateUpdateCategory(c *fiber.Ctx, userService userService.UserService, productService productService.ProductService) error {
	user := userService.GetUserModel(c)
	var params productDto.CategoryRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	category, hle := productService.SaveCategory(params, user)
	if hle.Status < 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  hle.Errors,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   productService.GetCategoryObject(category.ID),
	})
}

func ProductList(c *fiber.Ctx, userService userService.UserService, productService productService.ProductService) error {
	user := userService.GetUserModel(c)
	product := coreModel.Product{}
	queries := c.Queries()
	var andCondition []db.WhereCondition

	for key, value := range queries {
		if !utility.IsValidColumn(coreModel.ProductColumns, key) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -1,
				"error":  fmt.Sprintf("Invalid column %s", key),
			})
		}
		condition := db.WhereCondition{}
		condition.Key = key
		condition.Condition = "="
		condition.Value = value
		andCondition = append(andCondition, condition)
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
		andCondition = append(andCondition, db.WhereCondition{
			Key:       coreModel.ProductColumns.DistributorID,
			Condition: "=",
			Value:     distributorUser.DistributorID,
		})
	}

	list, err := product.FindAll(andCondition)
	if err != nil {
		// Return a 422 Unprocessable Entity status with the validation error
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -3,
			"error":  err,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   list,
	})

}

func CreateUpdateProduct(c *fiber.Ctx, userService userService.UserService, productService productService.ProductService) error {
	user := userService.GetUserModel(c)
	var params productDto.ProductRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	product, hle := productService.SaveProduct(params, user)
	if hle.Status < 0 {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  hle.Errors,
		})
	}

	return c.JSON(&fiber.Map{
		"status": 1,
		"data":   product,
	})
}
