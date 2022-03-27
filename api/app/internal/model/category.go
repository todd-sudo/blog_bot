package model

type Category struct {
	ID    uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name  string  `gorm:"type:varchar(255)" json:"name"`
	Posts *[]Post `json:"posts,omitempty"`
}
