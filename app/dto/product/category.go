package product_dto

type CategoryRequest struct {
	CategoryID       string `json:"category_id"`
	Code             string `json:"code" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description"`
	ParentCategoryID string `json:"parent_category_id"`
	Sequence         int    `json:"sequence" validate:"required"`
	Status           string `json:"status" validate:"required"`
}

type CategoryResponse struct {
	CategoryID       string            `json:"category_id"`
	Code             string            `json:"code" validate:"required"`
	Name             string            `json:"name" validate:"required"`
	Description      string            `json:"description"`
	ParentCategoryID string            `json:"parent_category_id"`
	ParentCategory   *CategoryResponse `json:"parent_category"`
	Sequence         int               `json:"sequence" validate:"required"`
	Status           string            `json:"status" validate:"required"`
}
