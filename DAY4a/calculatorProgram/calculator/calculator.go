package calculator

var c float64

func Add(a, b float64) float64 {
	c = a + b
	return c
}

func Subtract(a, b float64) float64 {
	c = a - b
	return c
}

func Multiply(a, b float64) float64 {
	c = a * b
	return c
}

func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	c = a / b
	return c
}

var res float64

func AddToLast(a float64) float64 {
	res = c + a
	return res
}
func SubtractFromLast(b float64) float64 {
	res = res - b
	return res
}
