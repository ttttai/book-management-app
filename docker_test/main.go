package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primary_key"`
	Name      string
	Email     string
	CreatedAt time.Time
}

func main() {
	dsn := "user:password@tcp(db:3306)/test_mysql"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	user := User{
		Name:  "test",
		Email: "test@example.com",
	}

	db.Create(&user)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping-pong-pong",
		})
	})
	r.Run()
}
