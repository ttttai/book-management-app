package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttttai/golang/usecases"
	"github.com/ttttai/golang/usecases/dto"
)

type IUserController interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetByName(c *gin.Context)
}

type UserController struct {
	userUsecase usecases.IUserUsecase
}

func NewUserController(userUsecase usecases.IUserUsecase) IUserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (uc *UserController) GetById(c *gin.Context) {
	var request dto.GetUserRequestParam

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

func (uc *UserController) Create(c *gin.Context) {
	var request dto.CreateUserRequestParam

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.userUsecase.Create(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (uc *UserController) Update(c *gin.Context) {
	var pathParam dto.UpdateUserRequestPathParam
	var bodyParam dto.UpdateUserRequestBodyParam

	if err := c.ShouldBindUri(&pathParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.userUsecase.Update(pathParam, bodyParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) Delete(c *gin.Context) {
	var request dto.DeleteUserRequestParam

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userUsecase.Delete(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (uc *UserController) GetByName(c *gin.Context) {
	var request dto.GetUserByNameRequestParam

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.userUsecase.GetByName(request.Name)
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
