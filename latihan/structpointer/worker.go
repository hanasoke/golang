package main

import "fmt"

type Worker struct {
	ID   int
	Name string
	GPA  float32
}

func (worker *Worker) graduate() {
	worker.Name = worker.Name + " B.Sc"
}

func main() {
	worker := Worker{1, "Hanas Bayu Pratama", 3.32}
	fmt.Println(worker.Name)
	worker.graduate()
	fmt.Println(worker.Name)
}
