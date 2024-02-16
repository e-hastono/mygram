package controllers

import (
	"net/http"
	"strconv"

	"github.com/e-hastono/mygram/helpers"
	"github.com/e-hastono/mygram/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotoInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

func CreatePhoto(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	var input PhotoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := helpers.ValidationErrorMessages(err.(validator.ValidationErrors))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  validationErrors,
		})
		return
	}

	sm := models.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.PhotoUrl,
		UserID:   user_id,
	}

	_, err = sm.SavePhoto()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "photo successfully created",
	})
}

func UpdatePhoto(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	var input PhotoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := helpers.ValidationErrorMessages(err.(validator.ValidationErrors))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  validationErrors,
		})
		return
	}

	sm := models.Photo{
		ID:       uint(photoId),
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.PhotoUrl,
		UserID:   user_id,
	}

	_, err = sm.UpdatePhoto()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "photo successfully updated",
	})
}

func DeletePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	sm := models.Photo{ID: uint(photoId)}
	err = sm.DeletePhoto()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "photo successfully deleted",
	})
}

func GetAllPhotos(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	photos, err := models.GetPhotosByUserID(user_id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": photos})
}

func GetOnePhoto(c *gin.Context) {
	photo_id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	photo, err := models.GetPhotoByPhotoIDUserID(uint(photo_id), user_id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": photo})
}
