package main

import "fmt"

// Calculator
type Calculator struct {
	a float64
	b float64
}

func (c Calculator) add() float64 {
	return c.a + c.b
}
func (c Calculator) sub() float64 {
	return c.a - c.b
}
func (c Calculator) mul() float64 {
	return c.a * c.b
}
func (c Calculator) div() float64 {
	return c.a / c.b
}

func main() {

	c := Calculator{2, 4}
	var op string
	fmt.Scanln(&op)
	switch op {
	case "+":
		fmt.Println("addition of 2 numbers", c.add())
	case "-":
		fmt.Println("Subtraction of two numbers", c.sub())
	case "*":
		fmt.Println("Multiplication of two numbers", c.mul())
	case "/":
		fmt.Println("Division of two numbers", c.div())
	default:
		fmt.Println("invalid operator")

	}
}
