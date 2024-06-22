package route

import (
	"example.com/awesome/dto"
	customValidator "example.com/awesome/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", func(c *gin.Context) {
			var registerUser dto.RegisterUserDto
			if err := c.ShouldBindJSON(&registerUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			validate := validator.New()

			if err := validate.Struct(registerUser); err != nil {
				errs, ok := customValidator.HandleValidationErrors(err)
				if ok {
					c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
					return
				}
			}
			c.JSON(200, gin.H{"message": "Registration successful", "data": registerUser})
		})
	}
}
