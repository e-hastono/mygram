package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"size:255;not null;unique" json:"username"`
	Email    string  `gorm:"size:255;not null;unique" json:"email"`
	Password string  `gorm:"size:255;not null;" json:"password"`
	Age      uint8   `gorm:"not null" json:"age"`
	Photos   []Photo `gorm:"foreignKey:UserID"`
}

type Photo struct {
	gorm.Model
	Title    string `gorm:"size:255;not null" json:"title"`
	Caption  string `gorm:"type:text;not null" json:"caption"`
	PhotoUrl string `gorm:"type:text;not null" json:"photo_url"`
	User     User
	UserID   uint
}

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"size:255;not null" json:"name"`
	SocialMediaUrl string `gorm:"type:text;not null" json:"social_media_url"`
	Title          string `gorm:"size:255;not null" json:"title"`
	Caption        string `gorm:"type:text;not null" json:"caption"`
	User           User
	UserID         uint `json:"user_id"`
}

type Comment struct {
	gorm.Model
	User    User
	UserID  uint
	Photo   Photo
	PhotoID uint
	Message string `gorm:"type:text;not null" json:"message"`
}
