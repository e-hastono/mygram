package main

import (
	"github.com/e-hastono/mygram/controllers"
	"github.com/e-hastono/mygram/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	public := r.Group("/api/v1")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run(":8080")

}
