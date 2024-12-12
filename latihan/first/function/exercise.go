package main

import "fmt"

func main() {

	// add := add(90, 30)
	// fmt.Println(add)
	// multipleby := multipleby(90, 9)
	// fmt.Println(multipleby)
	// dividedby := dividedby(60, 30)
	// fmt.Println(dividedby)
	// reduced := reduced(73, 40)
	// fmt.Println(reduced)

	keliling, _ := box(20, 3)
	fmt.Println(keliling)
}

func box(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

// func add(number, numberTwo int) int {
// 	return number + numberTwo
// }
// func multipleby(One, Two int) int {
// 	return One * Two
// }
// func dividedby(One, Two int) int {
// 	return One / Two
// }
// func reduced(One, Two int) int {
// 	return One - Two
// }
