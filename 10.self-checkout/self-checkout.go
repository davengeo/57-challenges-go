package main

import (
	"../shared"
	"fmt"
)

func main() {

	const tax_rate float64 = 0.055

	products:=[...]string{"1","2","3"}
	var prices [len(products)]int
	var amounts [len(products)]int

	for index, value := range products {
		prices[index]=shared.Get_number_with_question("Price of item " + value + "?")
		amounts[index]=shared.Get_number_with_question("Amount of item " + value + "?")
	}

	total_price:=0

	for index,_ := range products {
		total_price+=prices[index]*amounts[index]
	}

	fmt.Printf("Subtotal: %f\n", float64(total_price)*tax_rate)
	fmt.Printf("Taxes: %f\n", float64(total_price)*(1.0+tax_rate))
}
