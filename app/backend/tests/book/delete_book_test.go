package tests

import (
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

// 削除
func TestDeleteBookOK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, _ := infra.NewTestDB()
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	PrepareData(db)

	mockNdlApiRepository := new(mocks.MockNdlApiRepository)

	r := controllers.SetupTestRouter(db, mockNdlApiRepository)

	req, _ := http.NewRequest("DELETE", "/book/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var books []entities.Book
	var authors []entities.Author
	var subjects []entities.Subject
	var bookAuthors []entities.BookAuthor
	var bookSubjects []entities.BookSubject

	db.Find(&books)
	assert.Empty(t, books)

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
	assert.Empty(t, bookAuthors)

	db.Find(&bookSubjects)
	assert.Empty(t, bookSubjects)
}
