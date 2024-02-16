package controllers

import (
	"net/http"
	"strconv"

	"github.com/e-hastono/mygram/helpers"
	"github.com/e-hastono/mygram/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SocialMediaInput struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
	Title          string `json:"title"`
	Caption        string `json:"caption"`
}

func CreateSocialMedia(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	var input SocialMediaInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := helpers.ValidationErrorMessages(err.(validator.ValidationErrors))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  validationErrors,
		})
		return
	}

	sm := models.SocialMedia{
		Name:           input.Name,
		SocialMediaUrl: input.SocialMediaUrl,
		Title:          input.Title,
		Caption:        input.Caption,
		UserID:         user_id,
	}

	_, err = sm.SaveSocialMedia()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "social media successfully created",
	})
}

func UpdateSocialMedia(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	socialmediaId, err := strconv.Atoi(c.Param("socialmediaId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	var input SocialMediaInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := helpers.ValidationErrorMessages(err.(validator.ValidationErrors))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  validationErrors,
		})
		return
	}

	sm := models.SocialMedia{
		ID:             uint(socialmediaId),
		Name:           input.Name,
		SocialMediaUrl: input.SocialMediaUrl,
		Title:          input.Title,
		Caption:        input.Caption,
		UserID:         user_id,
	}

	_, err = sm.UpdateSocialMedia()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "social media successfully updated",
	})
}

func DeleteSocialMedia(c *gin.Context) {
	socialmediaId, err := strconv.Atoi(c.Param("socialmediaId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	sm := models.SocialMedia{ID: uint(socialmediaId)}
	err = sm.DeleteSocialMedia()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "social media successfully deleted",
	})
}

func GetAllSocialMedias(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	socialmedias, err := models.GetSocialMediasByUserID(user_id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": socialmedias})
}

func GetOneSocialMedia(c *gin.Context) {
	socialmedia_id, err := strconv.Atoi(c.Param("socialmedia_id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	socialmedia, err := models.GetSocialMediaBySocialMediaIDUserID(uint(socialmedia_id), user_id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": socialmedia})
}
