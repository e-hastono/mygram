package models

import (
	"errors"

	"github.com/e-hastono/mygram/database"
	"github.com/e-hastono/mygram/entities"
)

type Comment entities.Comment

func (com *Comment) SaveComment() (*Comment, error) {
	db := database.GetDB()

	err := db.Debug().Create(&com).Error
	if err != nil {
		return &Comment{}, err
	}

	return com, nil
}

func (com *Comment) UpdateComment() (*Comment, error) {
	db := database.GetDB()

	err := db.Debug().Model(&com).Updates(com).Error
	if err != nil {
		return &Comment{}, err
	}

	return com, nil
}

func (com *Comment) DeleteComment() error {
	db := database.GetDB()

	err := db.Debug().Delete(com).Error
	if err != nil {
		return err
	}

	return nil
}

func GetCommentsByUserID(uid uint) ([]Comment, error) {
	db := database.GetDB()

	var comments []Comment

	if err := db.Debug().Where("user_id = ?", uid).Find(&comments).Error; err != nil {
		return comments, errors.New("comments of user not found")
	}

	return comments, nil
}

func GetCommentByCommentIDUserID(comid uint, uid uint) (Comment, error) {
	var comment Comment

	db := database.GetDB()

	if err := db.Debug().Where("user_id = ?", uid).First(&comment, comid).Error; err != nil {
		return comment, errors.New("comment of user not found")
	}

	return comment, nil
}
