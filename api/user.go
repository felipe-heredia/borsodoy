package api

import (
	"borsodoy/radovid/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, _ := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes)
}

var users []*models.User

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, &users)
}

func CreateUser(c *gin.Context) {
	var newUserData models.CreateUser

	if err := c.BindJSON(&newUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      newUserData.Name,
		Email:     newUserData.Email,
		Password:  hashPassword(newUserData.Password),
		Balance:   0,
	}

	users = append(users, &newUser)
	c.IndentedJSON(http.StatusOK, newUser)
}
