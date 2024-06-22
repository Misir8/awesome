package handler

import (
	"example.com/awesome/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := service.GetUser()
	c.JSON(http.StatusOK, user)
}
