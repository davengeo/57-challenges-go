package main

import (
	"fmt"
)

func greeting(name string) {
	fmt.Println("Hello, " + name)
}

func main() {
	fmt.Print("What's your name?")
	var input string
	fmt.Scanln(&input)
	greeting(input)
}