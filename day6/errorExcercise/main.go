package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is a go tour excercise
type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number", float64(e))

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
