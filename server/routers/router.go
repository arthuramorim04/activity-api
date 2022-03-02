package routes

import (
	"github.com/arthuramorim04/go-activity-api.git/controllers"
	"github.com/arthuramorim04/go-activity-api.git/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoute(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		activity := main.Group("activity", middlewares.Auth())
		{
			activity.GET("/", controllers.ShowAllActivity)
			activity.GET("/:id", controllers.ShowActivity)
			activity.POST("/", controllers.CreateActivity)
			activity.PUT("/", controllers.EditActivity)
			activity.DELETE("/:id", controllers.DeleteActivity)
		}
		main.POST("/create-user", controllers.CreateUser)
		main.POST("/login", controllers.Login)
	}

	return router
}
