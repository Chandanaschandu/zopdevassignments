package calculatorMethods

import "fmt"

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
	if c.b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return c.a / c.b
}
