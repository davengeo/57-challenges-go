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