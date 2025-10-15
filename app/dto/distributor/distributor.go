package distributor_dto

import (
	userDto "infopack.co.in/offybox/app/dto/user"
)

type DistributorRequest struct {
	DistributorID  string `json:"distributor_id"`
	Name           string `json:"name" validate:"required"`
	Code           string `json:"code"`
	Mobile         string `json:"mobile" validate:"required"`
	Email          string `json:"email" validate:"required"`
	Status         string `json:"status" validate:"required"`
	PointOfContact string `json:"point_of_contact" validate:"required"`
}

type DistributorResponse struct {
	DistributorID  string             `json:"distributor_id"`
	Name           string             `json:"name"`
	Code           string             `json:"code"`
	Mobile         string             `json:"mobile"`
	Email          string             `json:"email"`
	Status         string             `json:"status"`
	PointOfContact string             `json:"point_of_contact"`
	CreatedAt      string             `json:"created_at"`
	UserDetail     userDto.UserObject `json:"user_detail"`
}

type DistributorUserRequest struct {
	DistributorID string `json:"distributor_id" validate:"required"`
	FirstName     string `json:"firstname" validate:"required"`
	LastName      string `json:"lastname" validate:"required"`
	Mobile        string `json:"mobile" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Status        string `json:"status" validate:"required"`
	RoleID        string `json:"role_id" validate:"required"`
	Password      string `json:"password"`
}
