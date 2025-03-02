package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/domain/services"
	"github.com/ttttai/golang/usecases"
	"github.com/ttttai/golang/usecases/dto"
	"gorm.io/gorm"
)

type IBookController interface {
	SearchBooks(c *gin.Context)
	GetBookInfoByBookId(c *gin.Context)
	CreateBookInfo(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
	UpdateBookStatus(c *gin.Context)
	GetBookInfo(c *gin.Context)
	GetGeminiResponse(c *gin.Context)
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
	var request dto.SearchBookRequestParam

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := bc.bookUsecase.SearchBooks(request.Title, request.MaxNum, request.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) GetBookInfoByBookId(c *gin.Context) {
	var request dto.GetBookInfoByBookIdRequestParam

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
			PublishDate:       request.Book.PublishDate,
			Price:             request.Book.Price,
			Status:            request.Book.Status,
			ReadingStartDate:  request.Book.ReadingStartDate,
			ReadingEndDate:    request.Book.ReadingEndDate,
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

func (bc *BookController) UpdateBook(c *gin.Context) {
	var request dto.UpdateBookRequestParam

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println(request)
	book := &entities.Book{
		ID:                request.ID,
		ISBN:              request.ISBN,
		TitleName:         request.TitleName,
		TitleNameKana:     request.TitleNameKana,
		PublisherName:     request.PublisherName,
		PublisherNameKana: request.PublisherNameKana,
		PublishDate:       request.PublishDate,
		Price:             request.Price,
		Status:            request.Status,
		ReadingStartDate:  request.ReadingStartDate,
		ReadingEndDate:    request.ReadingEndDate,
	}
	fmt.Println(book)

	res, err := bc.bookUsecase.UpdateBook(book)
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

func (bc *BookController) UpdateBookStatus(c *gin.Context) {
	var pathParam dto.UpdateBookStatusRequestPathParam
	var bodyParam dto.UpdateBookStatusRequestBodyParam

	if err := c.ShouldBindUri(&pathParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := bc.bookUsecase.UpdateBookStatus(pathParam.ID, bodyParam.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) GetBookInfo(c *gin.Context) {
	var request dto.GetBookInfoByStatusRequestParam

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	all_status := []int{services.BOOK_STATUS_NOT_PURCHASED, services.BOOK_STATUS_PURCHASED, services.BOOK_STATUS_READING, services.BOOK_STATUS_READ_COMPLETED}
	if len(request.Status) == 0 {
		request.Status = all_status
	}
	for _, statusItem := range request.Status {
		if !slices.Contains(all_status, statusItem) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "statusは0,1,2,3のいずれかを指定してください"})
			return
		}
	}

	res, err := bc.bookUsecase.GetBookInfo(request.Title, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) GetGeminiResponse(c *gin.Context) {
	var request dto.GetGeminiResponseRequestParam

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := bc.bookUsecase.GetGeminiResponse(request.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
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
