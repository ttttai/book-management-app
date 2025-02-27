package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// DBにbookInfoを登録
func TestCreateBookInfoOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	date := "2024-04-01"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	requestBody := dto.CreateBookInfoRequestParam{
		Book: dto.BookParam{
			ISBN:              1111111111111,
			TitleName:         "test",
			TitleNameKana:     "テスト",
			PublisherName:     "test社",
			PublisherNameKana: "テストシャ",
			PublishDate:       expectedDateTime,
			Price:             3000,
		},
		Authors: []dto.AuthorParam{
			{
				Name:     "test_name",
				NameKana: "テストネーム",
			},
			{
				Name:     "test_name2",
				NameKana: "テストネーム2",
			},
		},
		Subjects: []dto.SubjectParam{
			{
				SubjectName: "test_subject",
				SubjectKana: "テストサブジェクト",
			},
			{
				SubjectName: "test_subject2",
				SubjectKana: "テストサブジェクト2",
			},
		},
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/book", bytes.NewReader(requestBodyJson))
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

	var books []entities.Book
	var authors []entities.Author
	var subjects []entities.Subject
	var bookAuthors []entities.BookAuthor
	var bookSubjects []entities.BookSubject

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
		},
	}
	assert.Equal(t, expectedBooks, books)

	db.Find(&authors)
	expectedAuthors := []entities.Author{
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
	}
	assert.Equal(t, expectedAuthors, authors)

	db.Find(&subjects)
	expectedSubjects := []entities.Subject{
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
	}
	assert.Equal(t, expectedSubjects, subjects)

	db.Find(&bookAuthors)
	expectedBookAuthors := []entities.BookAuthor{
		{
			BookID:   1,
			AuthorID: 1,
		},
		{
			BookID:   1,
			AuthorID: 2,
		},
	}
	assert.Equal(t, expectedBookAuthors, bookAuthors)

	db.Find(&bookSubjects)
	expectedBookSubjects := []entities.BookSubject{
		{
			BookID:    1,
			SubjectID: 1,
		},
		{
			BookID:    1,
			SubjectID: 2,
		},
	}
	assert.Equal(t, expectedBookSubjects, bookSubjects)
}

// DBにすでに存在している本を登録
func TestCreateSameIsbnBookInfoOK(t *testing.T) {
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

	requestBody := dto.CreateBookInfoRequestParam{
		Book: dto.BookParam{
			ISBN:              1111111111111,
			TitleName:         "test",
			TitleNameKana:     "テスト",
			PublisherName:     "test社",
			PublisherNameKana: "テストシャ",
			PublishDate:       expectedDateTime,
			Price:             3000,
		},
		Authors: []dto.AuthorParam{
			{
				Name:     "test_name3",
				NameKana: "テストネーム3",
			},
			{
				Name:     "test_name4",
				NameKana: "テストネーム4",
			},
		},
		Subjects: []dto.SubjectParam{
			{
				SubjectName: "test_subject3",
				SubjectKana: "テストサブジェクト3",
			},
			{
				SubjectName: "test_subject4",
				SubjectKana: "テストサブジェクト4",
			},
		},
	}
	requestBodyJson, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/book", bytes.NewReader(requestBodyJson))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	fmt.Println(w.Body.String())

	var books []entities.Book
	var authors []entities.Author
	var subjects []entities.Subject
	var bookAuthors []entities.BookAuthor
	var bookSubjects []entities.BookSubject

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
		},
	}
	assert.Equal(t, expectedBooks, books)

	db.Find(&authors)
	expectedAuthors := []entities.Author{
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
	}
	assert.Equal(t, expectedAuthors, authors)

	db.Find(&subjects)
	expectedSubjects := []entities.Subject{
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
	}
	assert.Equal(t, expectedSubjects, subjects)

	db.Find(&bookAuthors)
	expectedBookAuthors := []entities.BookAuthor{
		{
			BookID:   1,
			AuthorID: 1,
		},
		{
			BookID:   1,
			AuthorID: 2,
		},
	}
	assert.Equal(t, expectedBookAuthors, bookAuthors)

	db.Find(&bookSubjects)
	expectedBookSubjects := []entities.BookSubject{
		{
			BookID:    1,
			SubjectID: 1,
		},
		{
			BookID:    1,
			SubjectID: 2,
		},
	}
	assert.Equal(t, expectedBookSubjects, bookSubjects)
}
