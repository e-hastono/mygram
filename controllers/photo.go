package controllers

import (
	"net/http"
	"strconv"

	"github.com/e-hastono/mygram/models"
	"github.com/e-hastono/mygram/utils/token"
	"github.com/gin-gonic/gin"
)

func GetAllPhotos(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
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
	photo_id, err := strconv.Atoi(c.Param("photo_id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	photo, err := models.GetPhotoByPhotoUserID(uint(photo_id), user_id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": photo})
}
