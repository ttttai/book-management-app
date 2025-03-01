package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repository_interfaces "github.com/ttttai/golang/domain/repositories"
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
	setupUserRoutes(r, userController)

	ndlApiRepository := repositories.NewNdlApiRepository()

	authorRepository := repositories.NewAuthorRepository(db)
	authorService := services.NewAuthorService(authorRepository)

	subjectRepository := repositories.NewSubjectRepository(db)
	subjectService := services.NewSubjectService(subjectRepository)

	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository, authorRepository, subjectRepository, ndlApiRepository)
	bookUsecase := usecases.NewBookUsecase(bookService, authorService, subjectService)
	bookController := NewBookController(bookUsecase)
	setupBookRoutes(r, bookController)

	return r
}

func SetupTestRouter(db *gorm.DB, mockNdlApiRepository repository_interfaces.INdlApiRepository) *gin.Engine {
	r := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NweUserUsecase(userRepository)
	userController := NewUserController(userUsecase)
	setupUserRoutes(r, userController)

	authorRepository := repositories.NewAuthorRepository(db)
	authorService := services.NewAuthorService(authorRepository)

	subjectRepository := repositories.NewSubjectRepository(db)
	subjectService := services.NewSubjectService(subjectRepository)

	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository, authorRepository, subjectRepository, mockNdlApiRepository)
	bookUsecase := usecases.NewBookUsecase(bookService, authorService, subjectService)
	bookController := NewBookController(bookUsecase)
	setupBookRoutes(r, bookController)

	return r
}

func setupUserRoutes(r *gin.Engine, userController IUserController) {
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
}

func setupBookRoutes(r *gin.Engine, bookController IBookController) {
	book := r.Group("/book")
	{
		book.GET("/search", bookController.SearchBooks)
		book.GET("/:id", bookController.GetBookInfoByBookId)
		book.GET("", bookController.GetBookInfo)
		book.POST("", bookController.CreateBookInfo)
		book.PUT("/:id", bookController.UpdateBook)
		book.PUT("/:id/status", bookController.UpdateBookStatus)
		book.DELETE("/:id", bookController.DeleteBook)
	}
}
