package main

import (
	"fmt"

	"github.com/ttttai/golang/controllers"
	"github.com/ttttai/golang/infra"
)

func main() {
	db, err := infra.NewDB()
	if err != nil {
		fmt.Println("failed to connect database")
		return
	}
	fmt.Println("Hello, World!")

	r := controllers.SetupRouter(db)
	r.Run()
}
