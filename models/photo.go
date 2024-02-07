package models

import (
	"errors"

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

func GetPhotosByUserID(uid uint) ([]Photo, error) {
	var photos []Photo

	if err := DB.Where("user_id = ?", uid).Find(&photos).Error; err != nil {
		return photos, errors.New("Photos of user not found")
	}

	return photos, nil
}

func GetPhotoByPhotoUserID(pid uint, uid uint) (Photo, error) {
	var photo Photo

	if err := DB.Where("user_id = ?", uid).First(&photo, pid).Error; err != nil {
		return photo, errors.New("Photo of user not found")
	}

	return photo, nil
}
