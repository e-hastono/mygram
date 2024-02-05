package models

import (
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"size:255;notnull" json:"name"`
	SocialMediaUrl string `gorm:"type:text;not null" json:"social_media_url"`
	Title          string `gorm:"size:255;not null" json:"title"`
	Caption        string `gorm:"type:text;not null" json:"caption"`
	User           User
	UserID         uint `json:"user_id"`
}
