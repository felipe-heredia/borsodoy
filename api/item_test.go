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

var itemID uuid.UUID

func Test_CreateItem(test *testing.T) {
	recorder := httptest.NewRecorder()

	requestBody := models.CreateItem{
		Name:      "Notebook",
		Price:     3000,
		ExpiresIn: 60,
		UserID:    user.ID,
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, "/item", strings.NewReader(string(requestBodyJson)))
	req.Header.Set("Authorization", accessToken)
	localRouter.ServeHTTP(recorder, req)

	var responseBody *models.Item
	json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	itemID = responseBody.ID

	assert.Equal(test, http.StatusOK, recorder.Code)
	assert.NotEqual(test, uuid.Nil, responseBody.ID)
	assert.Equal(test, requestBody.Name, responseBody.Name)
	assert.Equal(test, requestBody.Price, responseBody.Price)
	assert.Equal(test, requestBody.UserID, responseBody.UserID)

	shouldExpireAt := time.Now().Add(time.Duration(requestBody.ExpiresIn) * time.Minute)
	assert.GreaterOrEqual(test, shouldExpireAt, responseBody.ExpiredAt)
}

func Test_GetItemById(test *testing.T) {
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/item/%s", itemID.String())
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", accessToken)
	localRouter.ServeHTTP(recorder, req)

	var body models.Item
	json.Unmarshal(recorder.Body.Bytes(), &body)

	assert.Equal(test, http.StatusOK, recorder.Code)
	assert.Equal(test, itemID, body.ID)
}

func Test_GetItemById_shouldNotFindAnyItem(test *testing.T) {
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/item/%s", uuid.New().String())
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", accessToken)
	localRouter.ServeHTTP(recorder, req)

	var body struct{ Message string }
	json.Unmarshal(recorder.Body.Bytes(), &body)

	assert.Equal(test, http.StatusNotFound, recorder.Code)
	assert.Equal(test, "Item not found", body.Message)
}
