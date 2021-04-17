package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + ", A.Md.Kom"
}

func main() {
	student := Student{1, "Dafid Prasetyo", 3.95}

	fmt.Println("Nama sebelum gelar : ", student.Name)

	student.graduate()

	fmt.Println("Nama sesudah gelar :", student.Name)
}

// func graduate(student *Student) {

// 	student.Name = student.Name + ", A.Md.Kom"
// }
