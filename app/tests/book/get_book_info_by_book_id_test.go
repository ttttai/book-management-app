package tests

import (
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
)

// DBからbookIdでbookInfoを取得
func TestGetBookInfoByBookIdOK(t *testing.T) {
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

	req, _ := http.NewRequest("GET", "/book/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := entities.BookInfo{
		Book: entities.Book{
			ID:                1,
			ISBN:              1111111111111,
			TitleName:         "test",
			TitleNameKana:     "テスト",
			PublisherName:     "test社",
			PublisherNameKana: "テストシャ",
			PublishDate:       &expectedDateTime,
			Price:             3000,
		},
		Authors: []entities.Author{
			{
				ID:       1,
				Name:     "test_name",
				NameKana: "テストネーム",
			},
			{
				ID:       2,
				Name:     "test_name2",
				NameKana: "テストネーム2",
			},
		},
		Subjects: []entities.Subject{
			{
				ID:          1,
				SubjectName: "test_subject",
				SubjectKana: "テストサブジェクト",
			},
			{
				ID:          2,
				SubjectName: "test_subject2",
				SubjectKana: "テストサブジェクト2",
			},
		},
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())
}

// 存在しないIDを指定
func TestGetBookInfoByBookIdNG(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	PrepareData(db)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/book/1000", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := `{"error": "record not found"}`

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, expectedResponse, w.Body.String())

}
