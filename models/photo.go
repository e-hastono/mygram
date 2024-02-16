package models

import (
	"errors"

	"github.com/e-hastono/mygram/database"
	"github.com/e-hastono/mygram/entities"
)

type Photo entities.Photo

func (p *Photo) SavePhoto() (*Photo, error) {
	db := database.GetDB()

	err := db.Debug().Create(&p).Error
	if err != nil {
		return &Photo{}, err
	}

	return p, nil
}

func (p *Photo) UpdatePhoto() (*Photo, error) {
	db := database.GetDB()

	err := db.Debug().Model(&p).Updates(p).Error
	if err != nil {
		return &Photo{}, err
	}

	return p, nil
}

func (p *Photo) DeletePhoto() error {
	db := database.GetDB()

	err := db.Debug().Delete(p).Error
	if err != nil {
		return err
	}

	return nil
}

func GetPhotosByUserID(uid uint) ([]Photo, error) {
	db := database.GetDB()

	var photos []Photo

	if err := db.Debug().Where("user_id = ?", uid).Find(&photos).Error; err != nil {
		return photos, errors.New("photos of user not found")
	}

	return photos, nil
}

func GetPhotoByPhotoIDUserID(pid uint, uid uint) (Photo, error) {
	var photo Photo

	db := database.GetDB()

	if err := db.Debug().Where("user_id = ?", uid).First(&photo, pid).Error; err != nil {
		return photo, errors.New("photo of user not found")
	}

	return photo, nil
}
