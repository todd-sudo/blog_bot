package model

type User struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string  `gorm:"type:varchar(255)" json:"first_name,omitempty"`
	LastName  string  `gorm:"type:varchar(255)" json:"last_name,omitempty"`
	Username  string  `gorm:"type:varchar(255)" json:"username"`
	UserTGId  int64   `gorm:"type:varchar(255)" json:"user_tg_id"`
	Posts     *[]Post `json:"posts,omitempty"`
}
