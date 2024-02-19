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
		photo := apiGroup.Group("/photos")
		{
			photo.GET("/", controllers.GetAllPhoto)
			photo.GET("/:photoId", controllers.GetOnePhoto)
			photo.POST("/", controllers.CreatePhoto)
			photo.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
			photo.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
		}
		comment := apiGroup.Group("/comments")
		{
			comment.GET("/", controllers.GetAllComment)
			comment.GET("/:commentId", controllers.GetOneComment)
			comment.POST("/", controllers.CreateComment)
			comment.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
			comment.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
		}
		socialmedia := apiGroup.Group("/socialmedias")
		{
			socialmedia.GET("/", controllers.GetAllSocialMedia)
			socialmedia.GET("/:socialMediaId", controllers.GetOneSocialMedia)
			socialmedia.POST("/", controllers.CreateSocialMedia)
			socialmedia.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
			socialmedia.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
		}
	}

	return router
}
