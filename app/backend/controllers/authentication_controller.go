package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttttai/golang/domain/entities"
	"github.com/ttttai/golang/usecases"
	"github.com/ttttai/golang/usecases/dto"
)

type IAuthenticationController interface {
	GetByEmail(c *gin.Context)
	Create(c *gin.Context)
}

type AuthenticationController struct {
	authenticationUsecase usecases.IAuthenticationUsecase
}

func NewAuthenticationController(authenticationUsecase usecases.IAuthenticationUsecase) IAuthenticationController {
	return &AuthenticationController{
		authenticationUsecase: authenticationUsecase,
	}
}

func (ac *AuthenticationController) GetByEmail(c *gin.Context) {
	var request dto.GetByEmailRequestParam

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ac.authenticationUsecase.GetByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "存在しません。"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ac *AuthenticationController) Create(c *gin.Context) {
	var request dto.CreateAuthenticationRequestParam

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingAuthentication, err := ac.authenticationUsecase.GetByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if existingAuthentication != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "すでに登録されています。"})
		return
	}

	authentication := &entities.Authentication{
		Email:    request.Email,
		Password: request.Password,
	}

	res, err := ac.authenticationUsecase.Create(authentication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
