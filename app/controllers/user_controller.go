package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttttai/golang/usecases"
)

type IUserController interface {
	GetUser(c *gin.Context)
	Create(c *gin.Context)
}

type UserController struct {
	userUsecase usecases.IUserUsecase
}

type GetRequestParam struct {
	ID string `uri:"id" binding:"required"`
}

type CreateRequestParam struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func NewUserController(userUsecase usecases.IUserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (uc *UserController) GetById(c *gin.Context) {
	var request GetRequestParam

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.userUsecase.GetById(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}
