package models

import (
	"errors"

	"github.com/e-hastono/mygram/database"
	"github.com/e-hastono/mygram/entities"
)

type SocialMedia entities.SocialMedia

func (sm *SocialMedia) SaveSocialMedia() (*SocialMedia, error) {
	db := database.GetDB()

	err := db.Create(&sm).Error
	if err != nil {
		return &SocialMedia{}, err
	}

	return sm, nil
}

func (sm *SocialMedia) UpdateSocialMedia() (*SocialMedia, error) {
	db := database.GetDB()

	err := db.Debug().Model(&sm).Updates(sm).Error
	if err != nil {
		return &SocialMedia{}, err
	}

	return sm, nil
}

func (sm *SocialMedia) DeleteSocialMedia() error {
	db := database.GetDB()

	err := db.Delete(sm).Error
	if err != nil {
		return err
	}

	return nil
}

func GetSocialMediasByUserID(uid uint) ([]SocialMedia, error) {
	db := database.GetDB()

	var socialmedias []SocialMedia

	if err := db.Where("user_id = ?", uid).Find(&socialmedias).Error; err != nil {
		return socialmedias, errors.New("social medias of user not found")
	}

	return socialmedias, nil
}

func GetSocialMediaBySocialMediaIDUserID(smid uint, uid uint) (SocialMedia, error) {
	var socialmedia SocialMedia

	db := database.GetDB()

	if err := db.Where("user_id = ?", uid).First(&socialmedia, smid).Error; err != nil {
		return socialmedia, errors.New("social media of user not found")
	}

	return socialmedia, nil
}
