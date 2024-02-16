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
		protected.Use(middlewares.JwtAuthentication())

		// social media
		socialmedias := protected.Group("/socialmedia")
		{
			socialmedias.GET("/", controllers.GetAllSocialMedias)
			socialmedias.GET("/:socialmediaId", controllers.GetOneSocialMedia)

			socialmedias.POST("/", controllers.CreateSocialMedia)
			socialmedias.PUT("/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
			socialmedias.DELETE("/:socialmediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
		}

		// photos
		photos := protected.Group("/photos")
		{
			photos.GET("/", controllers.GetAllPhotos)
			photos.GET("/:photoId", controllers.GetOnePhoto)

			photos.POST("/", controllers.CreatePhoto)
			photos.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
			photos.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
		}
	}

	return r
}
