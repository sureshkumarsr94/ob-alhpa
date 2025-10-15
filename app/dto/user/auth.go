package user_dto

import masterDto "infopack.co.in/offybox/app/dto/master"

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password"  validate:"required"`
	Platform   string `json:"platform" validate:"required"`
}

type TokenDetail struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserType     string `json:"user_type"`
	TokenExpires int64  `json:"token_expires"`
}

type LoginResponse struct {
	Data   TokenDetail `json:"data"`
	Status int         `json:"status"`
}

type UserObject struct {
	UserID          string                 `json:"user_id"`
	AssociateUserId string                 `json:"associate_user_id"`
	Username        string                 `json:"username"`
	Firstname       string                 `json:"firstname"`
	Lastname        string                 `json:"lastname"`
	Mobile          string                 `json:"mobile"`
	UserType        string                 `json:"user_type"`
	Email           string                 `json:"email"`
	DistributorId   string                 `json:"distributor_id"`
	Role            masterDto.RoleResponse `json:"role"`
}

type ChangePasswordObject struct {
	UserId      string `json:"user_id" validate:"required_without=RequestId"`
	RequestId   string `json:"request_id" validate:"required_without=UserId"`
	OldPassword string `json:"old_password"`
	Password    string `json:"password" validate:"required"`
}
