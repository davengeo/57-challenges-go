package main

import (
	"fmt"
)

func main() {

	fmt.Println("What's your name?")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Hello, " + input)
}