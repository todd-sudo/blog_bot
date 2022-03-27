package dto

type CreateUserDTO struct {
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
	LastName  string `json:"last_name" form:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	UserTGId  int    `json:"user_tg_id"`
}
