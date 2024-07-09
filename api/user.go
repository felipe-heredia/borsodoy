package api

import (
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/internal/service"
	"borsodoy/radovid/pkg/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := service.GetUsers()

	c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUserData models.CreateUser

	if err := c.BindJSON(&newUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := service.CreateUser(newUserData)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			c.JSON(httpError.Status, gin.H{"message": httpError.Message})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserById(id)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			c.JSON(httpError.Status, gin.H{"message": httpError.Message})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, user)
}

func Protected(c *gin.Context) {
	email, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{"message": "Hello " + email.(string)})
}
