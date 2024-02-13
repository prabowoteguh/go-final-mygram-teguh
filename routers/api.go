package routers

import (
	"go-final-mygram/controllers"
	"go-final-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", controllers.UserRegister)
		auth.POST("/login", controllers.UserLogin)
	}

	apiGroup := router.Group("/api")
	{
		apiGroup.Use(middlewares.Authentication())
		photo := apiGroup.Group("/photo")
		{
			photo.GET("/", controllers.GetAllPhoto)
			photo.GET("/:photoId", controllers.GetOnePhoto)
			photo.POST("/", controllers.CreatePhoto)
			photo.PUT("/:photoId", controllers.UpdatePhoto)
			photo.DELETE("/:photoId", controllers.DeletePhoto)
		}
		comment := apiGroup.Group("/comment")
		{
			comment.GET("/", controllers.GetAllComment)
			comment.GET("/:commentId", controllers.GetOneComment)
			comment.POST("/", controllers.CreateComment)
			comment.PUT("/:commentId", controllers.UpdateComment)
			comment.DELETE("/:commentId", controllers.DeleteComment)
		}
		socialmedia := apiGroup.Group("/socialmedia")
		{
			socialmedia.GET("/", controllers.GetAllSocialMedia)
			socialmedia.GET("/:socialMediaId", controllers.GetOneSocialMedia)
			socialmedia.POST("/", controllers.CreateSocialMedia)
			socialmedia.PUT("/:socialMediaId", controllers.UpdateSocialMedia)
			socialmedia.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
		}
	}

	return router
}
