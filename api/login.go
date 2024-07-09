package api

import (
	"borsodoy/radovid/internal/service"
	"borsodoy/radovid/pkg/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginData service.LoginProps

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	responseBody, err := service.Login(loginData)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			c.JSON(httpError.Status, gin.H{"message": httpError.Message})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	c.IndentedJSON(http.StatusOK, responseBody)
}
