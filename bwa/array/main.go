package main

import "fmt"

func main() {

	/* How to fill the first array */
	var languages [5]string
	languages[0] = "Go"
	languages[1] = "PHP"
	languages[2] = "Javascript"
	languages[3] = "C++"
	languages[4] = "Ruby"

	/* How to fill the second array */
	programmingLanguages := [5]string{"Go", "PHP", "Javascript", "C++", "Ruby"}

	/*
		The colon in square brackets indicates the automatic calculation of the array length
		If the array is filled vertically, then the last element must still be marked with a comma
	*/
	programLanguages := [...]string{
		"Go",
		"PHP",
		"Javascript",
		"C++",
		"Ruby",
	}

	/* First Step */
	for index, lang := range programLanguages {
		fmt.Println("Index : ", index, " Language : ", lang)
	}
	/* Second Step */
	for i := 0; i < len(programLanguages); i++ {
		fmt.Println("Index : ", i, " Language : ", programLanguages[i])
	}

}
