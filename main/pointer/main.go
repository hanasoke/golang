package main

import "fmt"

func main() {
	numberA := 7
	numberB := &numberA

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)

	*numberB = 10

	fmt.Println(*numberB)
	fmt.Println(numberA)
}
