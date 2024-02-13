package controllers

import (
	"net/http"

	"github.com/e-hastono/mygram/helpers"
	"github.com/e-hastono/mygram/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      uint8  `json:"age" binding:"required,gt=8,lte=200"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := helpers.ValidationErrorMessages(err.(validator.ValidationErrors))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  validationErrors,
		})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password
	u.Age = input.Age

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "registration successful",
	})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := helpers.ValidationErrorMessages(err.(validator.ValidationErrors))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  validationErrors,
		})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	uid, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  "username or password is incorrect",
		})
		return
	}

	token, err := helpers.GenerateToken(uid)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}

func CurrentUser(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
