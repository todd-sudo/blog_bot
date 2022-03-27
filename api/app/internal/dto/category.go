package dto

type CreateCategoryDTO struct {
	Name string `json:"name" form:"name" binding:"required"`
}
