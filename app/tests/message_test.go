package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ttttai/golang/controllers"
	"github.com/ttttai/golang/infra"
)

type GetMessageResponse struct {
	Message string `json:"message"`
}

func TestGetMessageOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	r := controllers.SetupRouter(db)

	req, _ := http.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	message := GetMessageResponse{
		Message: "ping-pong-pong",
	}
	messageJson, _ := json.Marshal(message)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(messageJson), w.Body.String())
}
