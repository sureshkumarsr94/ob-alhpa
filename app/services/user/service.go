package user_service

import (
	"github.com/gofiber/fiber/v2"
	"infopack.co.in/offybox/app/dto"
	userDto "infopack.co.in/offybox/app/dto/user"
	coreModel "infopack.co.in/offybox/app/models/core"
)

// UserService defines the interface for user-related operations
type UserService interface {
	// ValidateCredentials validates the user's login credentials
	ValidateCredentials(request userDto.LoginRequest) (response userDto.LoginResponse, handle dto.HandleError)

	GetUserModel(c *fiber.Ctx) coreModel.User

	GetUserObject(userId string) userDto.UserObject
	// GenerateAuth generates authentication tokens for the user
	//GenerateAuth(users coreModel.User) (userDto.LoginResponse, error)
}

// userService is an implementation of UserService
type userService struct{}

// NewUserService returns a new instance of UserService
func NewUserService() UserService {
	return &userService{}
}
