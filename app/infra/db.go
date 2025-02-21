package infra

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := "user:password@tcp(db:3306)/test_database?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// タイムゾーンをJSTに設定
	loc, _ := time.LoadLocation("Asia/Tokyo")
	time.Local = loc

	return db, err
}
