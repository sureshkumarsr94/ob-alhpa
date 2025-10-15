package user_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null"
	"infopack.co.in/offybox/app/common/passwordutil"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/common/validator"
	userDto "infopack.co.in/offybox/app/dto/user"
	coreModel "infopack.co.in/offybox/app/models/core"
	userService "infopack.co.in/offybox/app/services/user"
	"net/http"
	"time"
)

// Login handles user authentication by validating the provided credentials
// Parameters:
// - c: *fiber.Ctx representing the request context
// - userService: userService.UserService for handling user-related operations
// Returns:
// - An error if there was an issue during the authentication process; otherwise, it returns a JSON response with the authentication details
func Login(c *fiber.Ctx, userService userService.UserService) error {
	// Initialize a LoginRequest DTO to hold the request parameters
	var params userDto.LoginRequest
	// Set the Platform from the request headers
	params.Platform = c.Get("X-Platform")

	// Parse and validate the request body into the params object
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		// Return a 422 Unprocessable Entity status with the validation error
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}

	// Call the userService to validate the provided credentials
	response, handle := userService.ValidateCredentials(params)
	if handle.Status < 0 {
		// Return a 422 Unprocessable Entity status with the service error
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": handle.Status,
			"error":  handle.Errors.Error(),
		})
	}

	// Return a JSON response with the authentication details
	return c.JSON(response)
}

func ChangePassword(c *fiber.Ctx) error {
	var params userDto.ChangePasswordObject
	if err := validator.ParseBodyAndValidate(c, &params); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -1,
			"error":  err,
		})
	}

	user := coreModel.User{}
	request := coreModel.UserCredentialRequest{}
	if params.RequestId != "" {
		request, _ = request.FindByPrimaryKey(params.RequestId)

		if request.ID == "" {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -2,
				"error":  "Request information not found",
			})
		}

		if utility.IsExpired(request.ExpireAt) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -3,
				"error":  "Reset request as expired",
			})
		}

		if request.Status == coreModel.CredentialRequestStatusUsed {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -4,
				"error":  "Request link expired",
			})
		}

		user = request.User
	}

	if user.ID == "" && params.UserId != "" {
		user, _ = user.FindByPrimaryKey(params.UserId)

		if user.ID == "" {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -5,
				"error":  "User information not found",
			})
		}

		if params.OldPassword != "" && user.Password != passwordutil.HashPassword(params.OldPassword) {
			return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": -6,
				"error":  "Invalid Password",
			})
		}
	}

	if !utility.IsValidPassword(params.Password) {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -7,
			"error":  "Password policy not met",
		})
	}

	user.Password = passwordutil.HashPassword(params.Password)
	user.LastPasswordChange = null.TimeFrom(time.Now())
	user, err := user.Save()

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": -7,
			"error":  err.Error(),
		})
	}

	if request.ID != "" && request.Status == coreModel.CredentialRequestStatusInitiated {
		request.Status = coreModel.CredentialRequestStatusUsed
		request, _ = request.Save()
	}

	return c.JSON(&fiber.Map{
		"status":  1,
		"message": "Password reset successfully done",
	})
}
