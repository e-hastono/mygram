package routers

import (
	"github.com/e-hastono/mygram/controllers"
	"github.com/e-hastono/mygram/middlewares"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api/v1")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)

		protected := public.Group("/user")
		protected.Use(middlewares.JwtAuthMiddleware())
		protected.GET("/", controllers.CurrentUser)

		// photos
		photos := protected.Group("/photos")
		{
			photos.GET("/", controllers.GetAllPhotos)
			photos.GET("/:photo_id", controllers.GetOnePhoto)
		}
	}

	return r
}
