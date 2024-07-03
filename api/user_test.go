package api

import (
	"borsodoy/radovid/internal/database"
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/router"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var localRouter *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	localRouter = router.SetupRouter()
	database.SetupTestDB()

	code := m.Run()

	os.Exit(code)
}

func Test_CreateUser(test *testing.T) {
	recorder := httptest.NewRecorder()

	exampleUser := models.CreateUser{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "securepassword",
	}
	userJson, _ := json.Marshal(exampleUser)

	req, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(string(userJson)))
	localRouter.ServeHTTP(recorder, req)

	var createdUser models.User
	json.Unmarshal(recorder.Body.Bytes(), &createdUser)

	assert.Equal(test, 200, recorder.Code)
	assert.Equal(test, exampleUser.Name, createdUser.Name)
	assert.Equal(test, uint32(0), createdUser.Balance)

	assert.NotEqual(test, uuid.Nil, createdUser.ID)
	assert.NotEqual(test, exampleUser.Password, createdUser.Password)
}

func Test_CreateUser_shouldErrDuplicatedKey(test *testing.T) {
	recorder := httptest.NewRecorder()

	exampleUser := models.CreateUser{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "securepassword",
	}
	userJson, _ := json.Marshal(exampleUser)

	req, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(string(userJson)))
	localRouter.ServeHTTP(recorder, req)

	var body struct{ Message string }
	json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Equal(test, 400, recorder.Code)
	assert.Equal(test, "Email already exists", body.Message)
}

func Test_CreateUser_shouldErrByJSONInvalid(test *testing.T) {
	recorder := httptest.NewRecorder()

	type CustomUser struct{ Name string }
	user := CustomUser{Name: "John Doe"}
	requestBody, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(string(requestBody)))
	localRouter.ServeHTTP(recorder, req)

	var body struct{ Message string }
	json.Unmarshal(recorder.Body.Bytes(), &body)
	fmt.Println(body)

	assert.Equal(test, 400, recorder.Code)
	assert.Equal(test, "Name, Email, and Password are required", body.Message)
}
