package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ttttai/golang/controllers"
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/infra"
	"github.com/ttttai/golang/tests/mocks"
	"github.com/ttttai/golang/usecases/dto"
)

// DBに登録
func TestCreateOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	requestBody := dto.CreateAuthenticationRequestParam{
		Email:    "test@test.com",
		Password: "test_password",
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/authentication/register", bytes.NewReader(requestBodyJson))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := entities.Authentication{
		ID:       1,
		Email:    "test@test.com",
		Password: "test_password",
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())

	var authentication entities.Authentication

	db.Find(&authentication)
	expectedAuthentication := entities.Authentication{
		ID:       1,
		Email:    "test@test.com",
		Password: "test_password",
	}
	assert.Equal(t, expectedAuthentication, authentication)
}
