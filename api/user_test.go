package api

import (
	"borsodoy/radovid/internal/database"
	"borsodoy/radovid/internal/middleware"
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/internal/service"
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

var (
	localRouter                        *gin.Engine
	user, bidderUser                   models.User
	johnDoeId, itemId                  uuid.UUID
	accessToken, bidderUserAccessToken string
)

func createGlobalUser() {
	recorder := httptest.NewRecorder()

	exampleUser := models.CreateUser{

		Name:     "Example User",
		Email:    "user@example.com",
		Password: "securepassword",
	}
	exampleUserJson, _ := json.Marshal(exampleUser)
	createUserReq, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(string(exampleUserJson)))
	localRouter.ServeHTTP(recorder, createUserReq)

	if recorder.Code != http.StatusOK {
		fmt.Printf("Failed to create user: %v\n", recorder.Body.String())
		return
	}

	json.Unmarshal(recorder.Body.Bytes(), &user)

  loginGlobalUser()
}

func createGlobalBidderUser() {
	recorder := httptest.NewRecorder()

	bidderUser := models.CreateUser{
		Name:     "Bidder User",
		Email:    "user@bidder.com",
		Password: "securepassword",
	}
	exampleUserJson, _ := json.Marshal(bidderUser)
	createUserReq, _ := http.NewRequest(http.MethodPost, "/user", strings.NewReader(string(exampleUserJson)))
	localRouter.ServeHTTP(recorder, createUserReq)

	if recorder.Code != http.StatusOK {
		fmt.Printf("Failed to create user: %v\n", recorder.Body.String())
		return
	}

	json.Unmarshal(recorder.Body.Bytes(), &bidderUser)

  loginGlobalBidderUser()
}

func loginGlobalUser() {
	recorder := httptest.NewRecorder()

	loginBody := service.LoginProps{
		Email:    "user@bidder.com",
		Password: "securepassword",
	}
	loginBodyJson, _ := json.Marshal(loginBody)
	loginReq, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(loginBodyJson)))
	localRouter.ServeHTTP(recorder, loginReq)

	if recorder.Code != http.StatusOK {
		fmt.Printf("Failed to login user: %v\n", recorder.Body.String())
		return
	}

	var loginBodyResponse struct {
		AccessToken string `json:"access_token"`
	}
	json.Unmarshal(recorder.Body.Bytes(), &loginBodyResponse)
	bidderUserAccessToken = loginBodyResponse.AccessToken
}

func loginGlobalBidderUser() {
	recorder := httptest.NewRecorder()

	loginBody := service.LoginProps{
		Email:    "user@example.com",
		Password: "securepassword",
	}
	loginBodyJson, _ := json.Marshal(loginBody)
	loginReq, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(loginBodyJson)))
	localRouter.ServeHTTP(recorder, loginReq)

	if recorder.Code != http.StatusOK {
		fmt.Printf("Failed to login user: %v\n", recorder.Body.String())
		return
	}

	var loginBodyResponse struct {
		AccessToken string `json:"access_token"`
	}
	json.Unmarshal(recorder.Body.Bytes(), &loginBodyResponse)
	accessToken = loginBodyResponse.AccessToken
}

func createGlobalItem() {
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

	itemId = responseBody.ID
}

func setupRouter() {
	gin.SetMode(gin.TestMode)
	localRouter = gin.Default()

	localRouter.GET("/users", GetUsers)
	localRouter.POST("/user", CreateUser)
	localRouter.GET("/user/:id", GetUserById)

	localRouter.POST("/login", Login)

	protected := localRouter.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/item", CreateItem)
		protected.GET("/item/:id", GetItemById)
		protected.POST("/bid", CreateBid)
		protected.DELETE("/bid/:id", WithdrawnBid)
	}
}

func TestMain(m *testing.M) {
	setupRouter()
	database.SetupTestDB()

	createGlobalUser()
  createGlobalBidderUser()
	createGlobalItem()

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
	johnDoeId = createdUser.ID

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

	assert.Equal(test, 400, recorder.Code)
	assert.Equal(test, "Name, Email, and Password are required", body.Message)
}

func Test_GetUsers(test *testing.T) {
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	localRouter.ServeHTTP(recorder, req)

	var body []models.User
	json.Unmarshal(recorder.Body.Bytes(), &body)

	assert.GreaterOrEqual(test, len(body), 1)
}

func Test_GetUserById(test *testing.T) {
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/user/%s", johnDoeId.String())
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	localRouter.ServeHTTP(recorder, req)

	var body models.User
	json.Unmarshal(recorder.Body.Bytes(), &body)
	fmt.Println(body)

	assert.Equal(test, http.StatusOK, recorder.Code)
	assert.Equal(test, johnDoeId, body.ID)
}

func Test_GetUserById_shouldNotFindAnyUser(test *testing.T) {
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/user/1", nil)
	localRouter.ServeHTTP(recorder, req)

	var body struct{ Message string }
	json.Unmarshal(recorder.Body.Bytes(), &body)

	assert.Equal(test, http.StatusNotFound, recorder.Code)
	assert.Equal(test, "User not found", body.Message)
}
