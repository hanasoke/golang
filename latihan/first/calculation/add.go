package calculation

func Add(number int, numberTwo int) int {
	return add(number, numberTwo)
}

func add(number int, numberTwo int) int {
	return number + numberTwo
}

func Min(numberOne int, numberTwo int) int {
	return min(numberOne, numberTwo)
}

func min(numberOne int, numberTwo int) int {
	return numberOne - numberTwo
}

func Kali(One int, Two int) int {
	return kali(One, Two)
}

func kali(One int, Two int) int {
	return One * Two
}

func Bagi(One int, Two int) int {
	return bagi(One, Two)
}

func bagi(One int, Two int) int {
	return One / Two
}

func Even(i int) int {
	return even(i)
}

func even(i int) int {
	return i % 2
}

func Odd(i int) int {
	return odd(i)
}

func odd(i int) int {
	return i % 1
}
