package main

import (
	"github.com/e-hastono/mygram/controllers"
	"github.com/e-hastono/mygram/models"
	"github.com/e-hastono/mygram/utils/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	public := r.Group("/api/v1")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)

		protected := public.Group("/admin")
		protected.Use(middleware.JwtAuthMiddleware())
		protected.GET("/user", controllers.CurrentUser)
	}

	r.Run(":8080")

}
