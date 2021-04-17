package main

import (
	"fmt"
	"hello-world/calculation"
)

func main() {

	fmt.Println("Program Perhitungan Aritmatik Dengan Bahasa Golang")

	add := calculation.Add(8, 9)
	multiplication := calculation.Multiplication(8, 9)

	fmt.Println("Hasil penjumlahan adalah :", add)
	fmt.Println("Hasil perkalian adalah :", multiplication)
}
