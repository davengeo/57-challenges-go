package main

import (
	"../shared"
	"fmt"
)

func main() {
	people:=shared.Get_number_with_question("How many people?")
	pizzas:=shared.Get_number_with_question("How many pizzas do you have?")
	slices:=shared.Get_number_with_question("How many slices per pizza?")

	fmt.Printf("%d people with %d pizzas.\n", people, pizzas)
	//the result of this is an int
	fmt.Printf("each person gets %d slices.\n", pizzas*slices/people)

	fmt.Printf("There is %d leftover pieces", pizzas*slices%people)
}
