package router

import (
	"borsodoy/radovid/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", api.GetUsers)
	router.POST("/user", api.CreateUser)
	router.GET("/user/:id", api.GetUserById)

  router.POST("/login", api.Login)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
