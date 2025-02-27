package infra

import (
	"fmt"
	"time"

	"github.com/ttttai/golang/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(options ...config.Options) (*gorm.DB, error) {
	config.LoadEnv(options...)
	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// タイムゾーンをJSTに設定
	loc, _ := time.LoadLocation("Asia/Tokyo")
	time.Local = loc

	return db, err
}

func NewTestDB() (*gorm.DB, error) {
	db, err := NewDB(config.WithMode("test"))
	if err != nil {
		return nil, err
	}

	if err := ResetTestDB(db); err != nil {
		return nil, err
	}

	return db, err
}

func ResetTestDB(db *gorm.DB) error {
	var tables []string

	if err := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE()").Scan(&tables).Error; err != nil {
		return err
	}

	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0;").Error; err != nil {
		return err
	}

	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE `%s`;", table)).Error; err != nil {
			return err
		}
	}

	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 1;").Error; err != nil {
		return err
	}

	return nil
}
