package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(test *testing.T) {
	router := SetupRouter()
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(test, 200, recorder.Code)

	expectedBody := gin.H{"message": "pong"}
	expectedJSON, err := json.Marshal(expectedBody)
	assert.NoError(test, err)

	assert.Equal(test, string(expectedJSON), recorder.Body.String())
}
