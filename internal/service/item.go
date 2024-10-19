package service

import (
	"borsodoy/radovid/internal/database"
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/pkg/utility"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CreateItem(data models.CreateItem) (*models.Item, error) {
	newItem := models.Item{
		ID:          uuid.New(),
		Name:        data.Name,
		Description: data.Description,
		ExpiredAt:   time.Now().Add(time.Duration(data.ExpiresIn) * time.Minute),
		Price:       data.Price,
		UserID:      data.UserID,
		ImageUrl:    data.ImageUrl,
	}

	if err := database.Database.Create(&newItem).Error; err != nil {
		return nil, &utility.HttpError{Message: "Internal Server Error", Status: http.StatusInternalServerError}
	}

	return &newItem, nil
}

func GetItemById(id string) (*models.Item, error) {
	var item *models.Item

	if err := database.Database.Preload("User").Preload("Bids", "withdrawn_at > ?", time.Now()).First(&item, "id = ?", id).Error; err != nil {
		return nil, &utility.HttpError{Message: "Item not found", Status: http.StatusNotFound}
	}

	return item, nil
}

func GetItems() ([]*models.Item, error) {
  var items []*models.Item

	database.Database.Preload("User").Preload("Bids").Find(&items, "expired_at > ?", time.Now())

	return items, nil
}
