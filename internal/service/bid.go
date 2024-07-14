package service

import (
	"borsodoy/radovid/internal/database"
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/pkg/utility"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CreateBid(data models.CreateBid) (*models.Bid, error) {
  item, _ := GetItemById(data.ItemID.String())

  if data.Amount < item.Price {
    return nil, &utility.HttpError{Message: "Bid amount smaller than item price", Status: http.StatusBadRequest}
  }

  if item.UserID == data.UserID {
    return nil, &utility.HttpError{Message: "You can't bid in your item", Status: http.StatusConflict}
  }

  newBid := models.Bid{
    ID: uuid.New(),
    Amount: data.Amount,
    ItemID: data.ItemID,
    UserID: data.UserID,
    WithdrawnAt: time.Now().Add(time.Duration(data.WithdrawnIn) * time.Minute),
  }

  if err := database.Database.Create(&newBid).Error; err != nil {
    return nil, &utility.HttpError{Message: "Internal Server Error", Status: http.StatusInternalServerError}
  }

  return &newBid, nil
}

func GetBidById(id string) (*models.Bid, error) {
  var bid *models.Bid

  if err := database.Database.First(&bid, "id = ?", id).Error; err != nil {
    return nil, &utility.HttpError{Message: "Bid not found", Status: http.StatusNotFound}
  }

  return bid, nil
}

func WithdrawnBid(id string) (bool, error) {
  bid, err := GetBidById(id)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
      return false, httpError
		}

    return false, &utility.HttpError{Message: "Internal Server Error", Status:http.StatusInternalServerError}
	}


  if err := database.Database.Model(&bid).Where("id = ?", id).Update("withdrawn_at", time.Now()).Error; err != nil {
    return false, &utility.HttpError{Message: "Internal Server Error", Status:http.StatusInternalServerError}
  }

  return true, nil
}
