package main

import (
	"fmt"
)

// 1.double the integer
func doubles(i int) int {
	return i * 2
}

// 2.print the input string
func greet(x string) string {
	return x

}

// perimeter of circle
func periCircle(r float32) float32 {
	const PI = 3.14
	return 2 * PI * r
}

//perimeter of squares

func periSquare(a int) int {
	return 4 * a

}
func cubeVolume(l, b, h int) int {
	return l * b * h

}

func periRectangle(l, b int) int {
	return 2 * (l + b)
}

func volumeSphere(ra float32) float32 {
	const PI = 3.14
	return 4 / 3 * (PI * ra * ra * ra)

}
func main() {

	i := doubles(4) //integer double
	fmt.Printf("Double the integer %d\n", i)

	x := greet("Hello world") //print input string
	fmt.Printf("Input string %s\n", x)

	pc := periCircle(10) //perimeter of circle
	fmt.Printf("perimeter of circle %g\n", pc)

	var ps int = periSquare(4) //perimeter of square
	fmt.Printf("Perimeter of square %d\n", ps)

	fmt.Printf("Perimeter of %d, %T\n", i, i)
	fmt.Printf("perimeter of %.2f", periCircle(10))

	var pr int = periRectangle(1, 1)
	fmt.Println(pr)

	var cv int = cubeVolume(2, 3, 4)
	fmt.Println(cv)
	var res float32 = volumeSphere(4)
	fmt.Println(res)

}
