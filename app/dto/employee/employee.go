package employee_dto

import masterDto "infopack.co.in/offybox/app/dto/master"

type EmployeeRequest struct {
	EmployeeID   string `json:"employee_id"`
	Name         string `json:"name" validate:"required"`
	Code         string `json:"code"`
	Mobile       string `json:"mobile" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Status       string `json:"status" validate:"required"`
	SupervisorID string `json:"supervisor_id"`
	RoleId       string `json:"role_id" validate:"required"`
	Password     string `json:"password"`
}

type EmployeeResponse struct {
	EmployeeID   string                 `json:"employee_id"`
	Name         string                 `json:"name"`
	Code         string                 `json:"code"`
	Mobile       string                 `json:"mobile"`
	Email        string                 `json:"email"`
	Status       string                 `json:"status"`
	SupervisorID string                 `json:"supervisor_id"`
	Supervisor   *EmployeeResponse      `json:"supervisor"`
	RoleID       string                 `json:"role_id"`
	Role         masterDto.RoleResponse `json:"role"`
	UserID       string                 `json:"user_id"`
}
