package main

import (
	"../shared"
	"fmt"
)
func main() {
	const factor float64 = 0.09290304
	length:=shared.Get_number_with_question("What is the length of the room in feet?")
	width:=shared.Get_number_with_question("What is the width of the room in feet?")
	fmt.Printf("You entered dimensions of %d feet by %d feet.\n", length, width)
	fmt.Printf("The area is \n%d square feet. \n%f square meters.\n", length*width, float64(length*width)*factor)
	fmt.Printf("=========================")
}

