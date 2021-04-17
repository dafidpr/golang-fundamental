package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println("Total dari variabel scores adalah : ", total)
	fmt.Println("======================================================")

	res, err := calculate(10, 2, "*")
	if err != nil {

		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println("Hasil Operasi adalah : ", res)
}

func sum(scores []int) int {
	var total int
	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}
	return total
}

func calculate(numOne int, numTwo int, operation string) (res int, err error) {

	if operation == "+" {
		res = numOne + numTwo
	} else if operation == "-" {
		res = numOne - numTwo
	} else if operation == "*" {
		res = numOne * numTwo
	} else if operation == "/" {
		res = numOne / numTwo
	} else {
		err = errors.New("Unknown Operation")
	}
	return
}
