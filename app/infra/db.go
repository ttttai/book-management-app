package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := "user:password@tcp(db:3306)/test_database?charset=utf8&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}
