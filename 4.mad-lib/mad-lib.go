package main

import (
	"fmt"
	"bufio"
	"os"
)


func read_line() string {
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	return string(input)
}

func main() {

	parts := [...]string {
		"noun",
		"verb",
		"adjective",
		"adverb",
	}

	var content [4]string

	for i, v := range parts {
		fmt.Println("Give me a " + v + "?")
		content[i] = read_line()
	}

	fmt.Printf("So... do you %s your %s %s %s? That's hilarious!", content[1], content[2], content[0], content[3])
}
