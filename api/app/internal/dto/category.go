package dto

type CreateCategoryDTO struct {
	Name   string `json:"name" form:"name" binding:"required"`
	UserID uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
