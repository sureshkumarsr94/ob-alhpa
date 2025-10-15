package master_dto

type RoleRequest struct {
	RoleID       string `json:"role_id"`
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	ParentRoleID string `json:"parent_role_id"`
	DataAccess   string `json:"data_access" validate:"required"`
	Status       string `json:"status" validate:"required"`
}

type RoleResponse struct {
	RoleID       string        `json:"role_id"`
	Code         string        `json:"code"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	ParentRoleID *string       `json:"parent_role_id"`
	ParentRole   *RoleResponse `json:"parent_role"`
	DataAccess   string        `json:"data_access"`
	Status       string        `json:"status"`
}
