package router

import (
	"borsodoy/radovid/api"
	"borsodoy/radovid/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", api.GetUsers)
	router.POST("/user", api.CreateUser)
	router.GET("/user/:id", api.GetUserById)

	router.POST("/login", api.Login)
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/protected", api.Protected)
	}

	return router
}
