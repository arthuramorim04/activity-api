package routes

import (
	"github.com/arthuramorim04/go-books-api.git/controllers"
	"github.com/arthuramorim04/go-books-api.git/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoute(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
		main.POST("/create-user", controllers.CreateUser)
		main.POST("/login", controllers.Login)
	}

	return router
}
