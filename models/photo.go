package models

import (
	"errors"

	"github.com/e-hastono/mygram/database"
	"github.com/e-hastono/mygram/entities"
)

type Photo entities.Photo

func GetPhotosByUserID(uid uint) ([]Photo, error) {
	db := database.GetDB()

	var photos []Photo

	if err := db.Where("user_id = ?", uid).Find(&photos).Error; err != nil {
		return photos, errors.New("Photos of user not found")
	}

	return photos, nil
}

func GetPhotoByPhotoIDUserID(pid uint, uid uint) (Photo, error) {
	var photo Photo

	db := database.GetDB()

	if err := db.Where("user_id = ?", uid).First(&photo, pid).Error; err != nil {
		return photo, errors.New("Photo of user not found")
	}

	return photo, nil
}
