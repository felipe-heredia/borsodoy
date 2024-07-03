package api

import (
	"borsodoy/radovid/internal/database"
	"borsodoy/radovid/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) string {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, _ := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes)
}

func GetUsers(c *gin.Context) {
	var users []*models.User

	database.Database.Find(&users)

	c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUserData models.CreateUser

	if err := c.BindJSON(&newUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUserData.Name == "" || newUserData.Email == "" || newUserData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name, Email, and Password are required"})
		return
	}

	newUser := models.User{
		ID:       uuid.New(),
		Name:     newUserData.Name,
		Email:    newUserData.Email,
		Password: hashPassword(newUserData.Password),
	}

	if err := database.Database.Create(&newUser).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newUser)
}
