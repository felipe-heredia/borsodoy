package api

import (
	"borsodoy/radovid/internal/database"
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/pkg/utility"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
		Password: utility.HashPassword(newUserData.Password),
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

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user *models.User

	if err := database.Database.First(&user, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
