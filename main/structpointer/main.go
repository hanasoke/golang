package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.T"
	// fmt.Println(student.Name)
}

func main() {
	student := Student{1, "Agung Setiawan", 3.27}
	fmt.Println(student.Name)
	student.graduate()
	fmt.Println(student.Name)
}
