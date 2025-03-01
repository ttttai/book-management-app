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
	"gorm.io/gorm"
)

func PrepareData(db *gorm.DB) {
	date := "2024-04-01"

	booksBefore := []entities.Book{
		{
			ID:                1,
			ISBN:              1111111111111,
			TitleName:         "test",
			TitleNameKana:     "テスト",
			PublisherName:     "test社",
			PublisherNameKana: "テストシャ",
			PublishDate:       &date,
			Price:             3000,
			Status:            0,
		},
	}
	db.Create(&booksBefore)

	authorsBefore := []entities.Author{
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
	db.Create(&authorsBefore)

	subjectsBefore := []entities.Subject{
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
	db.Create(&subjectsBefore)

	bookAuthorsBefore := []entities.BookAuthor{
		{
			BookID:   1,
			AuthorID: 1,
		},
		{
			BookID:   1,
			AuthorID: 2,
		},
	}
	db.Create(&bookAuthorsBefore)

	bookSubjectsBefore := []entities.BookSubject{
		{
			BookID:    1,
			SubjectID: 1,
		},
		{
			BookID:    1,
			SubjectID: 2,
		},
	}
	db.Create(&bookSubjectsBefore)
}

// 検索した本が重複なくDBに保存されるかをチェック
func TestSearchBooksOK(t *testing.T) {
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
	mockResponse := &[]entities.BookInfo{
		{
			Book: entities.Book{
				ISBN:              1111111111111,
				TitleName:         "test",
				TitleNameKana:     "テスト",
				PublisherName:     "test社",
				PublisherNameKana: "テストシャ",
				PublishDate:       &date,
				Price:             3000,
				Status:            0,
			},
			Authors: []entities.Author{
				{
					Name:     "test_name",
					NameKana: "テストネーム",
				},
				{
					Name:     "test_name2",
					NameKana: "テストネーム2",
				},
			},
			Subjects: []entities.Subject{
				{
					SubjectName: "test_subject",
					SubjectKana: "テストサブジェクト",
				},
				{
					SubjectName: "test_subject2",
					SubjectKana: "テストサブジェクト2",
				},
			},
		},
		{
			Book: entities.Book{
				ISBN:              2222222222222,
				TitleName:         "test2",
				TitleNameKana:     "テスト2",
				PublisherName:     "test2社",
				PublisherNameKana: "テストシャ2",
				PublishDate:       &date,
				Price:             5000,
				Status:            0,
			},
			Authors: []entities.Author{
				{
					Name:     "test_name",
					NameKana: "テストネーム",
				},
				{
					Name:     "test_name3",
					NameKana: "テストネーム3",
				},
			},
			Subjects: []entities.Subject{
				{
					SubjectName: "test_subject2",
					SubjectKana: "テストサブジェクト2",
				},
				{
					SubjectName: "test_subject3",
					SubjectKana: "テストサブジェクト3",
				},
			},
		},
	}

	mockNdlApiRepository.On("GetBooksFromNdlApi", "test", 2).Return(mockResponse, nil)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/book/search?title=test&maxNum=2", nil)
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
					ID:       1,
					Name:     "test_name",
					NameKana: "テストネーム",
				},
				{
					ID:       3,
					Name:     "test_name3",
					NameKana: "テストネーム3",
				},
			},
			Subjects: []entities.Subject{
				{
					ID:          2,
					SubjectName: "test_subject2",
					SubjectKana: "テストサブジェクト2",
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
			Status:            0,
		},
		{
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
		{
			ID:       3,
			Name:     "test_name3",
			NameKana: "テストネーム3",
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
		{
			ID:          3,
			SubjectName: "test_subject3",
			SubjectKana: "テストサブジェクト3",
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
		{
			BookID:   2,
			AuthorID: 1,
		},
		{
			BookID:   2,
			AuthorID: 3,
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
		{
			BookID:    2,
			SubjectID: 2,
		},
		{
			BookID:    2,
			SubjectID: 3,
		},
	}
	assert.Equal(t, expectedBookSubjects, bookSubjects)
}

// すでにDBに値が存在している場合の検索
func TestSearchBooksAlreadyRecordsExistOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	PrepareData(db)

	date := "2024-04-01"
	loc, _ := time.LoadLocation("Asia/Tokyo")
	dateTime, _ := time.ParseInLocation("2006-01-02", date, loc)
	expectedDateTime := dateTime.Format(time.RFC3339)
	mockResponse := &[]entities.BookInfo{
		{
			Book: entities.Book{
				ISBN:              1111111111111,
				TitleName:         "test",
				TitleNameKana:     "テスト",
				PublisherName:     "test社",
				PublisherNameKana: "テストシャ",
				PublishDate:       &date,
				Price:             3000,
				Status:            0,
			},
			Authors: []entities.Author{
				{
					Name:     "test_name",
					NameKana: "テストネーム",
				},
				{
					Name:     "test_name2",
					NameKana: "テストネーム2",
				},
			},
			Subjects: []entities.Subject{
				{
					SubjectName: "test_subject",
					SubjectKana: "テストサブジェクト",
				},
				{
					SubjectName: "test_subject2",
					SubjectKana: "テストサブジェクト2",
				},
			},
		},
		{
			Book: entities.Book{
				ISBN:              2222222222222,
				TitleName:         "test2",
				TitleNameKana:     "テスト2",
				PublisherName:     "test2社",
				PublisherNameKana: "テストシャ2",
				PublishDate:       &date,
				Price:             5000,
				Status:            0,
			},
			Authors: []entities.Author{
				{
					Name:     "test_name",
					NameKana: "テストネーム",
				},
				{
					Name:     "test_name3",
					NameKana: "テストネーム3",
				},
			},
			Subjects: []entities.Subject{
				{
					SubjectName: "test_subject2",
					SubjectKana: "テストサブジェクト2",
				},
				{
					SubjectName: "test_subject3",
					SubjectKana: "テストサブジェクト3",
				},
			},
		},
	}

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)
	mockNdlApiRepository.On("GetBooksFromNdlApi", "test", 2).Return(mockResponse, nil)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("GET", "/book/search?title=test&maxNum=2", nil)
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
					ID:       1,
					Name:     "test_name",
					NameKana: "テストネーム",
				},
				{
					ID:       3,
					Name:     "test_name3",
					NameKana: "テストネーム3",
				},
			},
			Subjects: []entities.Subject{
				{
					ID:          2,
					SubjectName: "test_subject2",
					SubjectKana: "テストサブジェクト2",
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
			Status:            0,
		},
		{
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
		{
			ID:       3,
			Name:     "test_name3",
			NameKana: "テストネーム3",
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
		{
			ID:          3,
			SubjectName: "test_subject3",
			SubjectKana: "テストサブジェクト3",
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
		{
			BookID:   2,
			AuthorID: 1,
		},
		{
			BookID:   2,
			AuthorID: 3,
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
		{
			BookID:    2,
			SubjectID: 2,
		},
		{
			BookID:    2,
			SubjectID: 3,
		},
	}
	assert.Equal(t, expectedBookSubjects, bookSubjects)
}
