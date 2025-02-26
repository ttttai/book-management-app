package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Mode struct {
	Mode string
}

type Options func(*Mode)

func WithMode(modeString string) Options {
	return func(mode *Mode) {
		mode.Mode = modeString
	}
}

func LoadEnv(options ...Options) {
	modeInitValue := &Mode{
		Mode: "prod",
	}
	for _, option := range options {
		option(modeInitValue)
	}

	envFile := ".env"
	if modeInitValue.Mode == "test" {
		envFile = "../.env.test"
	}
	fmt.Println(envFile)

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Warning: No %s file found\n", envFile)
	}
}

func GetDatabaseDSN() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	loc := "Asia%2FTokyo"

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		user, password, host, port, name, loc)
}
