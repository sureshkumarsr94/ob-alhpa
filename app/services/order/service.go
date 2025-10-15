package order_service

import (
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/dto"
	orderDto "infopack.co.in/offybox/app/dto/order"
	userDto "infopack.co.in/offybox/app/dto/user"
)

type OrderService interface {
	CreateUpdateOrder(tx *gorm.DB, request orderDto.OrderRequest, userObject userDto.UserObject) (outletCategoryID string, handle dto.HandleError)
}

// userService is an implementation of UserService
type orderService struct{}

// NewUserService returns a new instance of UserService
func NewUserService() OrderService {
	return &orderService{}
}
