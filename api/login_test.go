package api

import (
	"radovid/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Login(test *testing.T) {
	recorder := httptest.NewRecorder()

	requestBody := service.LoginProps{
		Email:    user.Email,
		Password: "securepassword",
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(requestBodyJson)))
	localRouter.ServeHTTP(recorder, req)

	var responseBody struct {
		AccessToken string `json:"access_token"`
	}
	json.Unmarshal(recorder.Body.Bytes(), &responseBody)

	assert.Equal(test, http.StatusOK, recorder.Code)
	assert.NotEmpty(test, responseBody.AccessToken)
}
