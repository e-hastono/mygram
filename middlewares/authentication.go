package middlewares

import (
	"net/http"

	"github.com/e-hastono/mygram/helpers"
	"github.com/gin-gonic/gin"
)

func JwtAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Next()
	}
}
