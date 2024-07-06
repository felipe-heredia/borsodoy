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
    return nil, &HttpError{ Message: "Name, Email, and Password are required", Status: http.StatusBadRequest }
	}

	newUser := models.User{
		ID:       uuid.New(),
		Name:     userData.Name,
		Email:    userData.Email,
		Password: utility.HashPassword(userData.Password),
	}

	if err := database.Database.Create(&newUser).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
      return nil, &HttpError{ Message: "Email already exists", Status: http.StatusBadRequest }
		}

    return nil, &HttpError{ Message: "Internal Server Error", Status: http.StatusInternalServerError }
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

  if err := database.Database.First(&user, "id = ?", id).Error; err != nil {
    return nil, &HttpError{ Message: "User not found", Status: http.StatusNotFound }
  }

  return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
  var user *models.User

  if err := database.Database.First(&user, "email = ?", email).Error; err != nil {
    return nil, &HttpError{ Message: "User not found", Status: http.StatusNotFound }
  }

  return user, nil
}
