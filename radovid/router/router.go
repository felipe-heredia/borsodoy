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
		protected.POST("/item", api.CreateItem)
		protected.GET("/item/:id", api.GetItemById)
		protected.POST("/bid", api.CreateBid)
		protected.DELETE("/bid/:id", api.WithdrawnBid)
	}

	return router
}
