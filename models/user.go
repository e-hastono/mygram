package models

import (
	"html"
	"strings"

	"github.com/e-hastono/mygram/database"
	"github.com/e-hastono/mygram/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User entities.User

func (u *User) SaveUser() (*User, error) {
	db := database.GetDB()

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (uint, error) {
	var err error

	u := User{}

	db := database.GetDB()

	err = db.Debug().Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return 0, err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, err
	}

	return u.ID, nil
}
