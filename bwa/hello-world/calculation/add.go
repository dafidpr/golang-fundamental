package calculation

/*
	Naming functions begin with a capital letter so that they can be accessed outside the package or can be accessed in a different package
*/
func Add(numberOne int, numberTwo int) int {

	/* @return numberOne + numberTwo */

	/* Example of calling func add() */
	return add(numberOne, numberTwo)
}

/*
	Naming functions beginning with lowercase letters can only be accessed in the same package
*/
func add(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}
