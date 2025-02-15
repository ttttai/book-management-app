package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	var message string = "hello23333"
	var num int = 10
	if num == 3 {
		fmt.Println(3)
	} else {
		fmt.Println(message)
	}
	fmt.Println(quote.Go())
}
