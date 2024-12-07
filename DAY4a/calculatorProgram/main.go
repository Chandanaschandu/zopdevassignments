package main

import (
	"fmt"
	"github.com/Chandanaschandu/DAY4a/calculatorProgram/calculator"
)

func main() {
	var a, b float64
	var operator string

	var add float64
	var sub float64
	var quit string
	for {

		fmt.Println("enter the two numbers")
		fmt.Scanln(&a, &b)

		fmt.Println("enter the operator")
		fmt.Scanln(&operator)

		switch operator {
		case "+":
			fmt.Println("Addition of two numbers", calculator.Add(a, b))
		case "-":
			fmt.Println("Subtraction of two numbers", calculator.Subtract(a, b))
		case "*":
			fmt.Println("Multiplication of two numbers", calculator.Multiply(a, b))
		case "/":
			fmt.Println("Division of two numbers", calculator.Divide(a, b))
		case "++":
			fmt.Println("Enter the number which should add to previous result")
			fmt.Scanln(&add)
			fmt.Println("adding to the last result", calculator.AddToLast(add))
		case "--":

			fmt.Println("Enter the number it should subtract from the previous result")
			fmt.Scanln(&sub)
			fmt.Println("Subracting  from  the last result", calculator.SubtractFromLast(sub))

		default:
			fmt.Println("Invalid operator")
		}

		fmt.Println("Do you want to continue? (type 'no' to quit, any other key to continue)")
		fmt.Scanln(&quit)
		if quit == "no" {
			fmt.Println("Exiting program.")
			break
		}

	}

}
