package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ttttai/golang/controllers"
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/infra"
	"github.com/ttttai/golang/tests/mocks"
	"github.com/ttttai/golang/usecases/dto"
)

// bookStatusを更新(0->1)
func TestUpdateBookStatusTo1OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	PrepareData(db)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	date := "2024-04-01"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	requestBody := dto.UpdateBookStatusRequestBodyParam{
		Status: 1,
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("PUT", "/book/status/1", bytes.NewReader(requestBodyJson))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := entities.Book{
		ID:                1,
		ISBN:              1111111111111,
		TitleName:         "test",
		TitleNameKana:     "テスト",
		PublisherName:     "test社",
		PublisherNameKana: "テストシャ",
		PublishDate:       &expectedDateTime,
		Price:             3000,
		Status:            1,
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())

	var books []entities.Book

	db.Find(&books)
	expectedBooks := []entities.Book{
		{
			ID:                1,
			ISBN:              1111111111111,
			TitleName:         "test",
			TitleNameKana:     "テスト",
			PublisherName:     "test社",
			PublisherNameKana: "テストシャ",
			PublishDate:       &expectedDateTime,
			Price:             3000,
			Status:            1,
		},
	}
	assert.Equal(t, expectedBooks, books)
}

// bookStatusを更新(0->2)
func TestUpdateBookStatusTo2OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	PrepareData(db)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	date := "2024-04-01"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	requestBody := dto.UpdateBookStatusRequestBodyParam{
		Status: 2,
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("PUT", "/book/status/1", bytes.NewReader(requestBodyJson))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := entities.Book{
		ID:                1,
		ISBN:              1111111111111,
		TitleName:         "test",
		TitleNameKana:     "テスト",
		PublisherName:     "test社",
		PublisherNameKana: "テストシャ",
		PublishDate:       &expectedDateTime,
		Price:             3000,
		Status:            2,
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())

	var books []entities.Book

	db.Find(&books)
	expectedBooks := []entities.Book{
		{
			ID:                1,
			ISBN:              1111111111111,
			TitleName:         "test",
			TitleNameKana:     "テスト",
			PublisherName:     "test社",
			PublisherNameKana: "テストシャ",
			PublishDate:       &expectedDateTime,
			Price:             3000,
			Status:            2,
		},
	}
	assert.Equal(t, expectedBooks, books)
}

// bookStatusを更新(0->3)
func TestUpdateBookStatusTo3OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	PrepareData(db)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	date := "2024-04-01"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	requestBody := dto.UpdateBookStatusRequestBodyParam{
		Status: 3,
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("PUT", "/book/status/1", bytes.NewReader(requestBodyJson))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := entities.Book{
		ID:                1,
		ISBN:              1111111111111,
		TitleName:         "test",
		TitleNameKana:     "テスト",
		PublisherName:     "test社",
		PublisherNameKana: "テストシャ",
		PublishDate:       &expectedDateTime,
		Price:             3000,
		Status:            3,
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())

	var books []entities.Book

	db.Find(&books)
	expectedBooks := []entities.Book{
		{
			ID:                1,
			ISBN:              1111111111111,
			TitleName:         "test",
			TitleNameKana:     "テスト",
			PublisherName:     "test社",
			PublisherNameKana: "テストシャ",
			PublishDate:       &expectedDateTime,
			Price:             3000,
			Status:            3,
		},
	}
	assert.Equal(t, expectedBooks, books)
}

// bookStatusに不適切な値を指定
func TestUpdateBookStatusInvalidStatusNG(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	PrepareData(db)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	date := "2024-04-01"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	requestBody := dto.UpdateBookStatusRequestBodyParam{
		Status: 4,
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("PUT", "/book/status/1", bytes.NewReader(requestBodyJson))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var books []entities.Book

	db.Find(&books)
	expectedBooks := []entities.Book{
		{
			ID:                1,
			ISBN:              1111111111111,
			TitleName:         "test",
			TitleNameKana:     "テスト",
			PublisherName:     "test社",
			PublisherNameKana: "テストシャ",
			PublishDate:       &expectedDateTime,
			Price:             3000,
			Status:            0,
		},
	}
	assert.Equal(t, expectedBooks, books)
}
