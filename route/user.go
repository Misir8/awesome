package route

import (
	"example.com/awesome/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/:id", handler.GetUser)
	}
}
