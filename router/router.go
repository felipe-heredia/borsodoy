package router

import (
	"radovid/api"
	"radovid/internal/middleware"

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
	router.GET("/items/", api.GetItems)
  router.GET("/item/:id", api.GetItemById)

	router.POST("/login", api.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/item", api.CreateItem)
		protected.POST("/bid", api.CreateBid)
		protected.DELETE("/bid/:id", api.WithdrawnBid)
	}

	return router
}
