package main

import "fmt"

func reverse(a []int) {
	var start = 0
	var stop = len(a) - 1

	for start < stop { //taking 2 pointers
		a[start], a[stop] = a[stop], a[start]
		start++
		stop--
	}

}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Before reversing an slice :%d\n", a)

	reverse(a)
	fmt.Printf("After reversing an slice :%d", a)
}
