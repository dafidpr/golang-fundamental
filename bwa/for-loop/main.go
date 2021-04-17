package main

import (
	"fmt"
)

func main() {

	/* For Loop style */
	for i := 0; i < 10; i++ {
		fmt.Println("Hallo")
	}

	/* While style */
	index := 1
	for index <= 10 {
		fmt.Println("Hallo Golang")
		index++
	}

	/* Foreach style */
	title := "Golang the best language"
	/*
		An undescore (_) is used if the variable is not used, then the alternative is to replace the variable with an undescore
	*/
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("Index : ", index, "Letter : ", string(letter))
		}
	}

	for index, letter := range title {
		letterString := string(letter)
		/* The first way to use if else */
		if letterString == "a" || letterString == "e" || letterString == "u" || letterString == "o" {
			fmt.Println("Index : ", index, "Letter : ", string(letter))
		}

		/* The second way is to use a switch case */
		switch letterString {
		case "a", "e", "u", "o":
			fmt.Println("Index : ", index, "Letter : ", string(letter))
		}
	}
}
