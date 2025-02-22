package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttttai/golang/usecases"
	"github.com/ttttai/golang/usecases/dto"
)

type IBookController interface {
	GetBooks(c *gin.Context)
}

type BookController struct {
	bookUsecase usecases.IBookUsecase
}

func NewBookController(bookUsecase usecases.IBookUsecase) IBookController {
	return &BookController{
		bookUsecase: bookUsecase,
	}
}

func (bc *BookController) GetBooks(c *gin.Context) {
	var request dto.GetBookRequestParam

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := bc.bookUsecase.GetBooks(request.Title, request.MaxNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
