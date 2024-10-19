package service

import (
	"borsodoy/radovid/internal/database"
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/pkg/utility"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(userData models.CreateUser) (*models.User, error) {
	if userData.Name == "" || userData.Email == "" || userData.Password == "" {
		return nil, &utility.HttpError{Message: "Name, Email, and Password are required", Status: http.StatusBadRequest}
	}

	newUser := models.User{
		ID:       uuid.New(),
		Name:     userData.Name,
		Email:    userData.Email,
		Password: utility.HashPassword(userData.Password),
	}

	if err := database.Database.Create(&newUser).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &utility.HttpError{Message: "Email already exists", Status: http.StatusBadRequest}
		}

		return nil, &utility.HttpError{Message: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return &newUser, nil
}

func GetUsers() []*models.User {
	var users []*models.User

	database.Database.Find(&users)

	return users
}

func GetUserById(id string) (*models.User, error) {
	var user *models.User

	if err := database.Database.Preload("Items").First(&user, "id = ?", id).Error; err != nil {
		return nil, &utility.HttpError{Message: "User not found", Status: http.StatusNotFound}
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user *models.User

	if err := database.Database.First(&user, "email = ?", email).Error; err != nil {
		return nil, &utility.HttpError{Message: "User not found", Status: http.StatusNotFound}
	}

	return user, nil
}
