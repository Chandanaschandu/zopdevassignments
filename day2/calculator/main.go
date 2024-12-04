package main

import "fmt"

func calculator(a float64, b float64, operator string) float64 {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}

}

func main() {
	fmt.Println("Enter two numbers a and b")

	var a, b float64
	var operator string

	fmt.Scanln(&a, &b)
	// message
	fmt.Println("Enter the operator")
	fmt.Scanln(&operator)

	fmt.Println("The result is ", calculator(a, b, operator))

}
