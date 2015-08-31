package main

import (
	"fmt"
)

func read_input() string {
	input := ""
	fmt.Scanln(&input)
	return input
}

func main() {
	fmt.Println("What is the input string?")
	input:=read_input()
	fmt.Printf("%s has %d characters", input, len(input))
}
