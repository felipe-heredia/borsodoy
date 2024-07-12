package api

import (
	"borsodoy/radovid/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var bid *models.Bid

func Test_CreateBid(test *testing.T) {
	recorder := httptest.NewRecorder()

	requestBody := models.CreateBid{
		Amount:      3100,
		WithdrawnIn: 60,
		ItemID:      itemId,
		UserID:      user.ID,
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, "/bid", strings.NewReader(string(requestBodyJson)))
	req.Header.Set("Authorization", accessToken)
	localRouter.ServeHTTP(recorder, req)

	json.Unmarshal(recorder.Body.Bytes(), &bid)

	assert.Equal(test, http.StatusOK, recorder.Code)
	assert.NotEqual(test, uuid.Nil, bid.ID)
	assert.Equal(test, bid.UserID, user.ID)
	assert.Equal(test, bid.ItemID, itemId)

	shouldExpireAt := time.Now().Add(time.Duration(requestBody.WithdrawnIn) * time.Minute)
	assert.GreaterOrEqual(test, shouldExpireAt, bid.WithdrawnAt)
}

func Test_WithdrawnBid(test *testing.T) {
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/bid/%s", bid.ID.String())
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Authorization", accessToken)
	localRouter.ServeHTTP(recorder, req)

	assert.Equal(test, http.StatusNoContent, recorder.Code)
}

func Test_CreateBid_shouldNotCreateWithAmountSmallerThanPrice(test *testing.T) {
	recorder := httptest.NewRecorder()

	requestBody := models.CreateBid{
		Amount:      100,
		WithdrawnIn: 60,
		ItemID:      itemId,
		UserID:      user.ID,
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, "/bid", strings.NewReader(string(requestBodyJson)))
	req.Header.Set("Authorization", accessToken)
	localRouter.ServeHTTP(recorder, req)

  var responseBody struct{Message string}
	json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.Equal(test, http.StatusBadRequest, recorder.Code)
  assert.Equal(test, "Bid amount smaller than item price", responseBody.Message)
}
