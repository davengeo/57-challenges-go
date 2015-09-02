package main

import (
	"../shared"
	"fmt"
)
func main() {
	length:=shared.Get_number_with_question("What is the length of the room in feet?")
	width:=shared.Get_number_with_question("What is the width of the room in feet?")
	fmt.Printf("%d %d \n", length, width)
}

