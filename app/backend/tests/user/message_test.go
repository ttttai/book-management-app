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
	"github.com/ttttai/golang/tests/mocks"
)

type GetMessageResponse struct {
	Message string `json:"message"`
}

func TestGetMessageOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()
	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

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
