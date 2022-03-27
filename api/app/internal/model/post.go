package model

type Post struct {
	ID         uint64   `gorm:"primary_key:auto_increment" json:"id"`
	Title      string   `gorm:"type:varchar(255)" json:"title"`
	Content    string   `gorm:"type:text" json:"content"`
	UserID     uint64   `gorm:"not null" json:"-"`
	User       User     `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CategoryID uint64   `gorm:"not null" json:"-"`
	Category   Category `gorm:"foreignkey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category"`
}