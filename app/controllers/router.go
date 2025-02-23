package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ttttai/golang/domain/services"
	"github.com/ttttai/golang/infra/repositories"
	"github.com/ttttai/golang/usecases"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NweUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	user := r.Group("/user")
	{
		user.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ping-pong-pong",
			})
		})
		user.GET("/search", userController.GetByName)
		user.GET("/:id", userController.GetById)
		user.POST("", userController.Create)
		user.PUT("/:id", userController.Update)
		user.DELETE("/:id", userController.Delete)
	}

	authorRepository := repositories.NewAuthorRepository(db)
	authorService := services.NewAuthorService(authorRepository)

	subjectRepository := repositories.NewSubjectRepository(db)
	subjectService := services.NewSubjectService(subjectRepository)

	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookUsecase := usecases.NewBookUsecase(bookService, authorService, subjectService)
	bookController := NewBookController(bookUsecase)

	book := r.Group("/book")
	{
		book.GET("/search", bookController.SearchBooks)
	}

	return r
}
