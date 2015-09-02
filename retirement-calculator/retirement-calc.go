package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"time"
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
	fmt.Println(pos)
	number,err=read_number()
	if (err!=nil) {
		get_number_with_question(pos)
	}
	return number
}

func main() {
	age:=get_number_with_question("What is your current age?")
	retire:=get_number_with_question("At what age would you like to retire?")
	if age>retire {
		fmt.Printf("Find a job, grandpa!")
		os.Exit(0)
	}
	fmt.Printf("You have %d years left until you can retire.\n", retire-age)
	now:=time.Now()
	fmt.Printf("It is %d, so you can retire in %d\n", now.Year(), now.Year()+retire-age)
}