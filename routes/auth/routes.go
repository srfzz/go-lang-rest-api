package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(server *gin.Engine) {
	authGroup := server.Group("/auth")
	{
		authGroup.POST("/signup", signup)
	}
}
