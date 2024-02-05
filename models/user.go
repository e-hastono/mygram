package models

import (
	"errors"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"size:255;not null;unique" json:"username"`
	Email    string  `gorm:"size:255;not null;unique" json:"email"`
	Password string  `gorm:"size:255;not null;" json:"password"`
	Age      uint8   `gorm:"type:tinyint unsigned; not null" json:"age"`
	Photos   []Photo `gorm:"foreignKey:UserID"`
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error
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

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return 0, err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, err
	}

	return u.ID, nil
}

func GetUserByID(uid uint) (User, error) {
	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found")
	}

	u.PrepareGive()

	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}
