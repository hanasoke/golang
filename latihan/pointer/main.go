package main

import "fmt"

func main() {
	// numberA := 10
	// numberB := &numberA
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 40
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// var numberA int = 9
	// var numberB *int = &numberA
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 20
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	number := 14
	fmt.Println("Alamat memory :", &number)
	fmt.Println("Nilai awal : ", number)
	change(&number, 100)
	fmt.Println("Nilai akhir : ", number)
	fmt.Println("Alamat Memory : ", &number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat memory :", &old)
	fmt.Println("Di dalam function :", *old)
}
