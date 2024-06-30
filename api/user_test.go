package api

import (
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/router"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(test *testing.T) {
	gin.SetMode(gin.TestMode)
	router := router.SetupRouter()
	recorder := httptest.NewRecorder()

	exampleUser := models.CreateUser{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "securepassword",
	}
	userJson, _ := json.Marshal(exampleUser)

	req, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(string(userJson)))
	router.ServeHTTP(recorder, req)

	var createdUser models.User
	json.Unmarshal(recorder.Body.Bytes(), &createdUser)

	assert.Equal(test, 200, recorder.Code)
	assert.Equal(test, exampleUser.Name, createdUser.Name)
	assert.NotEqual(test, uuid.Nil, createdUser.ID)
	assert.NotEqual(test, exampleUser.Password, createdUser.Password)
}
