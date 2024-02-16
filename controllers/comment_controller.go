package controllers

import (
	"net/http"
	"strconv"

	"github.com/e-hastono/mygram/helpers"
	"github.com/e-hastono/mygram/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentInput struct {
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func CreateComment(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	var input CommentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := helpers.ValidationErrorMessages(err.(validator.ValidationErrors))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  validationErrors,
		})
		return
	}

	sm := models.Comment{
		PhotoID: input.PhotoID,
		Message: input.Message,
		UserID:  user_id,
	}

	_, err = sm.SaveComment()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "comment successfully created",
	})
}

func UpdateComment(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	var input CommentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErrors := helpers.ValidationErrorMessages(err.(validator.ValidationErrors))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  validationErrors,
		})
		return
	}

	sm := models.Comment{
		ID:      uint(commentId),
		PhotoID: input.PhotoID,
		Message: input.Message,
		UserID:  user_id,
	}

	_, err = sm.UpdateComment()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "comment successfully updated",
	})
}

func DeleteComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	sm := models.Comment{ID: uint(commentId)}
	err = sm.DeleteComment()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "comment successfully deleted",
	})
}

func GetAllComments(c *gin.Context) {
	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "failure",
			"error":  err.Error(),
		})
		return
	}

	comments, err := models.GetCommentsByUserID(user_id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": comments})
}

func GetOneComment(c *gin.Context) {
	comment_id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user_id, err := helpers.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	comment, err := models.GetCommentByCommentIDUserID(uint(comment_id), user_id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": comment})
}
