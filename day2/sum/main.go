package main

import "fmt"

func sum(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	return sum

}
func main() {
	fmt.Println("Enter the number to sum")
	var m int
	fmt.Scanln(&m)
	fmt.Printf("Sum of the  number is %d", sum(m))
}
