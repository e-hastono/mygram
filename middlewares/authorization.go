package middlewares

import (
	"net/http"
	"strconv"

	"github.com/e-hastono/mygram/database"
	"github.com/e-hastono/mygram/helpers"
	"github.com/e-hastono/mygram/models"
	"github.com/gin-gonic/gin"
)

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialmediaId, err := strconv.Atoi(c.Param("socialmediaId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		userID, err := helpers.ExtractTokenID(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		SocialMedia := models.SocialMedia{}

		err = db.Debug().Select("user_id").First(&SocialMedia, uint(socialmediaId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		if SocialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failure",
				"error":  "unauthorized to access this data",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		userID, err := helpers.ExtractTokenID(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		Photo := models.Photo{}

		err = db.Debug().Select("user_id").First(&Photo, uint(photoId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		if Photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failure",
				"error":  "unauthorized to access this data",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		userID, err := helpers.ExtractTokenID(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		Comment := models.Comment{}

		err = db.Debug().Select("user_id").First(&Comment, uint(commentId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": "failure",
				"error":  err.Error(),
			})
			return
		}

		if Comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "failure",
				"error":  "unauthorized to access this data",
			})
			return
		}

		c.Next()
	}
}
