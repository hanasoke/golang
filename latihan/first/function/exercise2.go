package main

import (
	"errors"
	"fmt"
)

func calculator(a, b int, operator string) (int, error) {
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

func main() {
	result, err := calculator(13, 32, "+")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
	result, err = calculator(9, 14, "-")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
	result, err = calculator(12, 4, "*")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
	result, err = calculator(14, 2, "/")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}
