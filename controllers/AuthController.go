package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mira-moonbeam/go-auth-be/middlewares"
	"github.com/mira-moonbeam/go-auth-be/services"
)

func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	{
		public := authGroup.Group("/public")
		{
			public.POST("/register", services.Register)
			public.POST("/login", services.Login)
		}

		// Endpoints requiring token
		protected := authGroup.Group("/protected")
		{
			protected.Use(middlewares.JwtAuthMiddleware())
			protected.GET("/user", services.GetCurrentUser)
		}
	}
}
