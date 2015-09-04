package main

import (
	"../shared"
	"fmt"
)

type basket_item struct {
	name string
	price float64
	amount int
}

func main() {

	const tax_rate float64 = 0.055

	var basket [3]basket_item
	total_price:=0.0

	basket[0].name="1"
	basket[1].name="2"
	basket[2].name="3"

	for _, value := range basket {
		value.price=shared.Get_float_with_question("Price of item " + value.name + "?")
		value.amount=shared.Get_number_with_question("Amount of item " + value.name + "?")
		total_price+=value.price*float64(value.amount)
	}

	fmt.Printf("Subtotal: %f\n", total_price)
	fmt.Printf("Taxes: %f\n", total_price*tax_rate)
	fmt.Printf("Total: %f\n", total_price*(1.0 + tax_rate))
}
