package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/usecases"
	"github.com/ttttai/golang/usecases/dto"
	"gorm.io/gorm"
)

type IBookController interface {
	SearchBooks(c *gin.Context)
	GetBookInfoByBookId(c *gin.Context)
	CreateBookInfo(c *gin.Context)
	DeleteBook(c *gin.Context)
}

type BookController struct {
	bookUsecase usecases.IBookUsecase
}

func NewBookController(bookUsecase usecases.IBookUsecase) IBookController {
	return &BookController{
		bookUsecase: bookUsecase,
	}
}

func (bc *BookController) SearchBooks(c *gin.Context) {
	var request dto.GetBookRequestParam

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := bc.bookUsecase.SearchBooks(request.Title, request.MaxNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) GetBookInfoByBookId(c *gin.Context) {
	var request dto.GetBookInfoRequestParam

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := bc.bookUsecase.GetBookInfoByBookId(request.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) CreateBookInfo(c *gin.Context) {
	var request dto.CreateBookInfoRequestParam

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	bookInfo := &entities.BookInfo{
		Book: entities.Book{
			ISBN:              request.Book.ISBN,
			TitleName:         request.Book.TitleName,
			TitleNameKana:     request.Book.TitleNameKana,
			PublisherName:     request.Book.PublisherName,
			PublisherNameKana: request.Book.PublisherNameKana,
			PublishDate:       &request.Book.PublishDate,
			Price:             request.Book.Price,
		},
		Authors:  convertAuthors(request.Authors),
		Subjects: convertSubjects(request.Subjects),
	}

	res, err := bc.bookUsecase.CreateBookInfo(bookInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	var request dto.DeleteBookRequestParam

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := bc.bookUsecase.DeleteBook(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func convertAuthors(authorsParam []dto.AuthorParam) []entities.Author {
	var authors []entities.Author
	for _, author := range authorsParam {
		authors = append(authors, entities.Author{
			Name:     author.Name,
			NameKana: author.NameKana,
		})
	}
	return authors
}

func convertSubjects(subjectParam []dto.SubjectParam) []entities.Subject {
	var subjects []entities.Subject
	for _, subject := range subjectParam {
		subjects = append(subjects, entities.Subject{
			SubjectName: subject.SubjectName,
			SubjectKana: subject.SubjectKana,
		})
	}
	return subjects
}
