package entities

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string  `gorm:"size:255;not null;unique" json:"username"`
	Email     string  `gorm:"size:255;not null;unique" json:"email"`
	Password  string  `gorm:"size:255;not null;" json:"password"`
	Age       uint8   `gorm:"not null" json:"age"`
	Photos    []Photo `gorm:"foreignKey:UserID"`
}

type SocialMedia struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string `gorm:"size:255;not null" json:"name"`
	SocialMediaUrl string `gorm:"type:text;not null" json:"social_media_url"`
	User           User
	UserID         uint `json:"user_id"`
}

type Photo struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `gorm:"size:255;not null" json:"title"`
	Caption   string `gorm:"type:text;not null" json:"caption"`
	PhotoUrl  string `gorm:"type:text;not null" json:"photo_url"`
	User      User
	UserID    uint
}

type Comment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	UserID    uint
	Photo     Photo
	PhotoID   uint
	Message   string `gorm:"type:text;not null" json:"message"`
}
