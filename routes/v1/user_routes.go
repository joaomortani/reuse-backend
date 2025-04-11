package v1

import (
	handlers "reuse-api/handlers/user"
	"reuse-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, handler *handlers.UserHandler) {
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.POST("/refresh", handler.RefreshToken)
	}
	user := r.Group("/api/v1/user")
	user.Use(middlewares.AuthMiddleware())
	{
		user.GET("/profile", handler.Profile)
	}
}
