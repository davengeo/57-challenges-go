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

	basket[0].name="1"
	basket[1].name="2"
	basket[2].name="3"

	for index, value := range basket {
		basket[index].price=shared.Get_float_with_question("Price of item " + value.name + "?")
		basket[index].amount=shared.Get_number_with_question("Amount of item " + value.name + "?")
	}

	total_price:=0.0

	for _, value := range basket {
		total_price+=value.price*float64(value.amount)
	}

	fmt.Printf("Subtotal: %f\n", total_price*tax_rate)
	fmt.Printf("Taxes: %f\n", total_price*(1.0 + tax_rate))
}
