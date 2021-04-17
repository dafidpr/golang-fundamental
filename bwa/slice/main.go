package main

import "fmt"

func main() {

	/* The first way of filling the slice */
	var gamingConsoles []string
	gamingConsoles = append(gamingConsoles, "PlayStation 2")
	gamingConsoles = append(gamingConsoles, "PlayStation 4")
	gamingConsoles = append(gamingConsoles, "Xbox One")

	/* The second way of filling the slice */
	gamingConsole := []string{"PlayStation 2", "PlayStation 4"}

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}
