package shared

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func read_number() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	return strconv.Atoi(string(input))
}

func Get_number_with_question(pos string) int {
	var (
		number int
		err error
	)
	fmt.Println(pos)
	number,err=read_number()
	if (err!=nil) {
		Get_number_with_question(pos)
	}
	return number
}

func read_float() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	return strconv.ParseFloat(string(input), 10)
}

func Get_float_with_question(pos string) float64 {
	var (
		number float64
		err error
	)
	fmt.Println(pos)
	number,err=read_float()
	if (err!=nil) {
		Get_float_with_question(pos)
	}
	return number
}