package order_controller

import (
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/common/validator"
	"infopack.co.in/offybox/app/database"
	orderDto "infopack.co.in/offybox/app/dto/order"
	outletService "infopack.co.in/offybox/app/services/order"
	userService "infopack.co.in/offybox/app/services/user"
	"log"
	"net/http"
)

func CreateUpdateOrder(c *fiber.Ctx, orderService outletService.OrderService, userService userService.UserService) error {
	var params orderDto.OrderRequest
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}
	userModel := userService.GetUserModel(c)
	userObject := userService.GetUserObject(userModel.ID)
	log.Println("userObject :: ", userObject)
	tx := database.MysqlDB.Begin()
	orderId, hle := orderService.CreateUpdateOrder(tx, params, userObject)
	if hle.Status < 0 {
		tx.Rollback()
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": hle.Status,
			"error":  hle.Errors.Error(),
		})
	}
	return c.JSON(&fiber.Map{
		"status": 1,
		"result": &fiber.Map{
			"order_id": orderId,
		},
	})
}
