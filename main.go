package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mira-moonbeam/go-auth-be/controllers"
	"github.com/mira-moonbeam/go-auth-be/models"
)

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// Connect DB
	models.ConnectDatabase()

	// Create a new Gin router
	router := gin.Default()
	router.Use(CORS())

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!")
	})

	// Inject Routes
	controllers.AuthRoutes(router)

	// Run the server on port 8081
	err := router.Run(":8081")
	if err != nil {
		println("SERVER CANNOT START")
		return
	}
}
