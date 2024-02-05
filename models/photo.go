package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `gorm:"size:255;not null" json:"title"`
	Caption  string `gorm:"type:text;not null" json:"caption"`
	PhotoUrl string `gorm:"type:text;not null" json:"photo_url"`
	User     User
	UserID   uint
}
