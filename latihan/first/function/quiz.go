package main

import (
	"errors"
	"fmt"
)

// func sum(scores []int) int {
// 	total := 0
// 	for _, score := range scores {
// 		total += score
// 	}
// 	return total
// }

func main() {
	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)

	// scores := []int{10, 5, 8, 9, 7}
	// total := sum(scores)
	// fmt.Printf("Total score: %d\n", total)

	// result, err := calculate(10, 2, "+")
	// result, err := calculate(10, 2, "-")
	// result, err := calculate(10, 2, "*")
	// result, err := calculate(10, 2, "/")
	// result, err := calculate(10, 2, "=")

	result, err := calculate(10, 2, "+")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
	result, err = calculate(10, 2, "-")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
	result, err = calculate(10, 2, "*")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
	result, err = calculate(10, 2, "/")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
	result, err = calculate(10, 2, "=")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Division by zero is not allowed")
		}
		return a / b, nil
	default:
		return 0, errors.New("Unsupported operator")
	}
}
