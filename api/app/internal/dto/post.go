package dto

type PostCreateDTO struct {
	Title      string `json:"title" form:"title" binding:"required"`
	Content    string `json:"content" form:"content" binding:"required"`
	UserID     uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	CategoryID uint64 `json:"category_id,omitempty" form:"category_id,omitempty"`
}
