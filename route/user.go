package route

import (
	"example.com/awesome/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userHandler := handler.NewUserHandler()
	userGroup := router.Group("/user")
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.GET("/", userHandler.GetUsers)
	}
}
