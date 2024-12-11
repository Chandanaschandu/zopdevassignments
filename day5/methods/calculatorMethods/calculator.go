package calculatorMethods

import "fmt"

type Calculator struct {
}

func (c Calculator) Add(A float64, B float64) float64 {
	return A + B
}
func (c Calculator) Sub(A float64, B float64) float64 {
	return A - B
}
func (c Calculator) Mul(A float64, B float64) float64 {
	return A * B
}
func (c Calculator) Div(A float64, B float64) float64 {
	if B == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return A / B
}
