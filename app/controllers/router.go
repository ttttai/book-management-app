package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(d *gorm.DB) *gin.Engine {
	r := gin.Default()

	todo := r.Group("/todo")
	{
		todo.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "test",
			})
		})
	}

	return r
}
