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

// DBからbookInfoを全件取得
func TestGetBookInfoOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	date := "2024-04-01"
	PrepareData(db)
	bookBefore := []entities.Book{
		{
			ISBN:              2222222222222,
			TitleName:         "test2",
			TitleNameKana:     "テスト2",
			PublisherName:     "test2社",
			PublisherNameKana: "テストシャ2",
			PublishDate:       &date,
			Price:             5000,
			Status:            0,
		},
	}
	db.Create(bookBefore)

	authorBefore := []entities.Author{
		{
			Name:     "test_name3",
			NameKana: "テストネーム3",
		},
	}
	db.Create(&authorBefore)

	subjectBefore := []entities.Subject{
		{
			SubjectName: "test_subject3",
			SubjectKana: "テストサブジェクト3",
		},
	}
	db.Create(&subjectBefore)

	bookAuthorsBefore := []entities.BookAuthor{
		{
			BookID:   2,
			AuthorID: 2,
		},
		{
			BookID:   2,
			AuthorID: 3,
		},
	}
	db.Create(&bookAuthorsBefore)

	bookSubjectsBefore := []entities.BookSubject{
		{
			BookID:    2,
			SubjectID: 1,
		},
		{
			BookID:    2,
			SubjectID: 3,
		},
	}
	db.Create(&bookSubjectsBefore)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/book", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := []entities.BookInfo{
		{
			Book: entities.Book{
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
		},
		{
			Book: entities.Book{
				ID:                2,
				ISBN:              2222222222222,
				TitleName:         "test2",
				TitleNameKana:     "テスト2",
				PublisherName:     "test2社",
				PublisherNameKana: "テストシャ2",
				PublishDate:       &expectedDateTime,
				Price:             5000,
				Status:            0,
			},
			Authors: []entities.Author{
				{
					ID:       2,
					Name:     "test_name2",
					NameKana: "テストネーム2",
				},
				{
					ID:       3,
					Name:     "test_name3",
					NameKana: "テストネーム3",
				},
			},
			Subjects: []entities.Subject{
				{
					ID:          1,
					SubjectName: "test_subject",
					SubjectKana: "テストサブジェクト",
				},
				{
					ID:          3,
					SubjectName: "test_subject3",
					SubjectKana: "テストサブジェクト3",
				},
			},
		},
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())
}

// DBからbookInfoをtitleで取得
func TestGetBookInfoByTitleOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	date := "2024-04-01"
	PrepareData(db)
	bookBefore := []entities.Book{
		{
			ISBN:              2222222222222,
			TitleName:         "test2",
			TitleNameKana:     "テスト2",
			PublisherName:     "test2社",
			PublisherNameKana: "テストシャ2",
			PublishDate:       &date,
			Price:             5000,
			Status:            0,
		},
	}
	db.Create(bookBefore)

	authorBefore := []entities.Author{
		{
			Name:     "test_name3",
			NameKana: "テストネーム3",
		},
	}
	db.Create(&authorBefore)

	subjectBefore := []entities.Subject{
		{
			SubjectName: "test_subject3",
			SubjectKana: "テストサブジェクト3",
		},
	}
	db.Create(&subjectBefore)

	bookAuthorsBefore := []entities.BookAuthor{
		{
			BookID:   2,
			AuthorID: 2,
		},
		{
			BookID:   2,
			AuthorID: 3,
		},
	}
	db.Create(&bookAuthorsBefore)

	bookSubjectsBefore := []entities.BookSubject{
		{
			BookID:    2,
			SubjectID: 1,
		},
		{
			BookID:    2,
			SubjectID: 3,
		},
	}
	db.Create(&bookSubjectsBefore)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/book?title=test2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := []entities.BookInfo{
		{
			Book: entities.Book{
				ID:                2,
				ISBN:              2222222222222,
				TitleName:         "test2",
				TitleNameKana:     "テスト2",
				PublisherName:     "test2社",
				PublisherNameKana: "テストシャ2",
				PublishDate:       &expectedDateTime,
				Price:             5000,
				Status:            0,
			},
			Authors: []entities.Author{
				{
					ID:       2,
					Name:     "test_name2",
					NameKana: "テストネーム2",
				},
				{
					ID:       3,
					Name:     "test_name3",
					NameKana: "テストネーム3",
				},
			},
			Subjects: []entities.Subject{
				{
					ID:          1,
					SubjectName: "test_subject",
					SubjectKana: "テストサブジェクト",
				},
				{
					ID:          3,
					SubjectName: "test_subject3",
					SubjectKana: "テストサブジェクト3",
				},
			},
		},
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())
}

// DBからbookInfoをstatusで取得
func TestGetBookInfoByStatusOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	date := "2024-04-01"
	PrepareData(db)
	bookBefore := []entities.Book{
		{
			ISBN:              2222222222222,
			TitleName:         "test2",
			TitleNameKana:     "テスト2",
			PublisherName:     "test2社",
			PublisherNameKana: "テストシャ2",
			PublishDate:       &date,
			Price:             5000,
			Status:            1,
		},
	}
	db.Create(bookBefore)

	authorBefore := []entities.Author{
		{
			Name:     "test_name3",
			NameKana: "テストネーム3",
		},
	}
	db.Create(&authorBefore)

	subjectBefore := []entities.Subject{
		{
			SubjectName: "test_subject3",
			SubjectKana: "テストサブジェクト3",
		},
	}
	db.Create(&subjectBefore)

	bookAuthorsBefore := []entities.BookAuthor{
		{
			BookID:   2,
			AuthorID: 2,
		},
		{
			BookID:   2,
			AuthorID: 3,
		},
	}
	db.Create(&bookAuthorsBefore)

	bookSubjectsBefore := []entities.BookSubject{
		{
			BookID:    2,
			SubjectID: 1,
		},
		{
			BookID:    2,
			SubjectID: 3,
		},
	}
	db.Create(&bookSubjectsBefore)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/book?status=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := []entities.BookInfo{
		{
			Book: entities.Book{
				ID:                2,
				ISBN:              2222222222222,
				TitleName:         "test2",
				TitleNameKana:     "テスト2",
				PublisherName:     "test2社",
				PublisherNameKana: "テストシャ2",
				PublishDate:       &expectedDateTime,
				Price:             5000,
				Status:            1,
			},
			Authors: []entities.Author{
				{
					ID:       2,
					Name:     "test_name2",
					NameKana: "テストネーム2",
				},
				{
					ID:       3,
					Name:     "test_name3",
					NameKana: "テストネーム3",
				},
			},
			Subjects: []entities.Subject{
				{
					ID:          1,
					SubjectName: "test_subject",
					SubjectKana: "テストサブジェクト",
				},
				{
					ID:          3,
					SubjectName: "test_subject3",
					SubjectKana: "テストサブジェクト3",
				},
			},
		},
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())
}

// DBからbookInfoを複数のstatusで取得
func TestGetBookInfoByStatusesOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	date := "2024-04-01"
	PrepareData(db)
	bookBefore := []entities.Book{
		{
			ISBN:              2222222222222,
			TitleName:         "test2",
			TitleNameKana:     "テスト2",
			PublisherName:     "test2社",
			PublisherNameKana: "テストシャ2",
			PublishDate:       &date,
			Price:             5000,
			Status:            1,
		},
		{
			ISBN:              3333333333333,
			TitleName:         "test3",
			TitleNameKana:     "テスト3",
			PublisherName:     "test3社",
			PublisherNameKana: "テストシャ3",
			PublishDate:       &date,
			Price:             10000,
			Status:            2,
		},
	}
	db.Create(bookBefore)

	authorBefore := []entities.Author{
		{
			Name:     "test_name3",
			NameKana: "テストネーム3",
		},
	}
	db.Create(&authorBefore)

	subjectBefore := []entities.Subject{
		{
			SubjectName: "test_subject3",
			SubjectKana: "テストサブジェクト3",
		},
	}
	db.Create(&subjectBefore)

	bookAuthorsBefore := []entities.BookAuthor{
		{
			BookID:   2,
			AuthorID: 2,
		},
		{
			BookID:   2,
			AuthorID: 3,
		},
		{
			BookID:   3,
			AuthorID: 1,
		},
	}
	db.Create(&bookAuthorsBefore)

	bookSubjectsBefore := []entities.BookSubject{
		{
			BookID:    2,
			SubjectID: 1,
		},
		{
			BookID:    2,
			SubjectID: 3,
		},
		{
			BookID:    3,
			SubjectID: 1,
		},
	}
	db.Create(&bookSubjectsBefore)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/book?status=1&status=2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := []entities.BookInfo{
		{
			Book: entities.Book{
				ID:                2,
				ISBN:              2222222222222,
				TitleName:         "test2",
				TitleNameKana:     "テスト2",
				PublisherName:     "test2社",
				PublisherNameKana: "テストシャ2",
				PublishDate:       &expectedDateTime,
				Price:             5000,
				Status:            1,
			},
			Authors: []entities.Author{
				{
					ID:       2,
					Name:     "test_name2",
					NameKana: "テストネーム2",
				},
				{
					ID:       3,
					Name:     "test_name3",
					NameKana: "テストネーム3",
				},
			},
			Subjects: []entities.Subject{
				{
					ID:          1,
					SubjectName: "test_subject",
					SubjectKana: "テストサブジェクト",
				},
				{
					ID:          3,
					SubjectName: "test_subject3",
					SubjectKana: "テストサブジェクト3",
				},
			},
		},
		{
			Book: entities.Book{
				ID:                3,
				ISBN:              3333333333333,
				TitleName:         "test3",
				TitleNameKana:     "テスト3",
				PublisherName:     "test3社",
				PublisherNameKana: "テストシャ3",
				PublishDate:       &expectedDateTime,
				Price:             10000,
				Status:            2,
			},
			Authors: []entities.Author{
				{
					ID:       1,
					Name:     "test_name",
					NameKana: "テストネーム",
				},
			},
			Subjects: []entities.Subject{
				{
					ID:          1,
					SubjectName: "test_subject",
					SubjectKana: "テストサブジェクト",
				},
			},
		},
	}
	expectedResponseJson, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedResponseJson), w.Body.String())
}

// 存在しないstatusを指定
func TestGetBookInfoByInvalidStatusNG(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	PrepareData(db)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/book?status=100", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expectedResponse := `{"error": "statusは0,1,2,3のいずれかを指定してください"}`

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, expectedResponse, w.Body.String())

}
