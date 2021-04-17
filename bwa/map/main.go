package main

import "fmt"

func main() {

	/* Declaration myMap */
	var myMap map[string]int

	/* Initialize myMap */
	myMap = map[string]int{}

	/* The first way to fill in the map */
	myMap["ruby"] = 9
	myMap["go"] = 9
	myMap["javascript"] = 9

	/* The second way is map declaration and initialization */
	myLang := map[string]string{
		"ruby":       "is beautyful language",
		"go":         "is super fast",
		"javascript": "hype",
	}

	for key, val := range myLang {
		fmt.Println("Key : ", key, " Value : ", val)
	}
	fmt.Println("==============")

	/* Delete map */
	delete(myLang, "ruby")

	for key, val := range myLang {
		fmt.Println("Key : ", key, " Value : ", val)
	}

	/*
		Check whether the value from the map is available or not
		The @value variable will contain or return the contents of the map.
		Whereas the variable @isAvailable will return a boolean value,
		if the key is available it will return true, otherwise it will return false
	*/
	value, isAvailable := myLang["python"]
	fmt.Println(isAvailable)
	fmt.Println(value)
}
