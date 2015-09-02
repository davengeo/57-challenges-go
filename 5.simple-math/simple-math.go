package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func read_number() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	return strconv.Atoi(string(input))
}


func get_number_with_question(pos string) int {
	var (
		number int
		err error
	)
	loop := 0
	for loop==0 {
		fmt.Println(pos)
		number,err=read_number()
		if (err==nil) {
			loop = 1
		}
	}
	return number
}

func main() {
	first := get_number_with_question("Whay is the first number?")
	second := get_number_with_question("What is the second number?")
	if (second==0) {
		fmt.Println("Zero is not allowed as second number.")
		os.Exit(1)
	}
	fmt.Printf("%d+%d = %d\n", first, second, first+second)
	fmt.Printf("%d-%d = %d\n", first, second, first-second)
	fmt.Printf("%d*%d = %d\n", first, second, first*second)
	fmt.Printf("%d/%d = %f\n", first, second, first/second)
}
