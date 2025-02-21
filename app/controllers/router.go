package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttttai/golang/infra/repositories"
	"github.com/ttttai/golang/usecases"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NweUserUsecase(userRepository)
	UserController := NewUserController(userUsecase)

	todo := r.Group("/user")
	{
		todo.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ping-pong-pong",
			})
		})
		todo.GET("/:id", UserController.GetById)
		todo.POST("", UserController.Create)
		todo.PUT("/:id", UserController.Update)
		todo.DELETE("/:id", UserController.Delete)
	}

	return r
}
