package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User    User
	UserID  uint
	Photo   Photo
	PhotoID uint
	Message string `gorm:"type:text;not null" json:"message"`
}
