package tests

import (
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
)

// DBからEmailを指定して取得
func TestGetByEmailOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	authentication := []entities.Authentication{
		{
			ID:       1,
			Email:    "test@test.com",
			Password: "test_password",
		},
	}
	db.Create(&authentication)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/authentication/search?email=test@test.com", nil)
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
}

// DBに存在しない場合
func TestGetByEmailNotExistOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/authentication/search?email=test@test.com", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := `{"error": "存在しません。"}`

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, expectedResponse, w.Body.String())
}
