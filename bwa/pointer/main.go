package main

import "fmt"

func main() {

	numberA := 5
	/* numberB contains the memory address where the contents of the numberA variable are stored */
	numberB := &numberA

	fmt.Println(numberA)
	fmt.Println(numberB)
	/* return to the beginning as before */
	fmt.Println(*numberB)

	/* assign value to numerB */
	*numberB = 10

	fmt.Println(*numberB)
	fmt.Println(numberA)

	/* Pointer 2 */
	var numA int = 5
	var numB *int = &numA

	fmt.Println(numA)
	fmt.Println(numB)

	numA = 20

	fmt.Println(numA)
	fmt.Println(*numB)
	fmt.Println(numB)

	fmt.Println("===============================")

	number := 5
	fmt.Println("Nilai awal : ", number)

	/* send memory address parameters */
	change(&number, 100)

	fmt.Println("Nilai akhir : ", number)

}

func change(oldVal *int, newVal int) {
	*oldVal = newVal
}
