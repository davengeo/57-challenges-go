package main

import (
	"../shared"
	"fmt"
)


func round_up(length int, width int) int {
	const gallon2feet float64 = 350

	var result float64
	result = float64(length)*float64(width)/ gallon2feet
	if result > float64(int(result)) {
		result++
	}
	return int(result)
}

func main() {

	length:=shared.Get_number_with_question("What is the length?")
	width:=shared.Get_number_with_question("What is the width?")

	fmt.Printf("You will need to purchase %d gallons of paint to cover %d square feet.\n", round_up(length, width), length*width)

}