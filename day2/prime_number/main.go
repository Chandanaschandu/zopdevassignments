package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	} else if n == 2 || n == 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%i+1 == 0 {
			return false
		}
	}

	return true
}

func main() {
	//To check number is prime or not
	var n int
	fmt.Println("enter the number")

	fmt.Scanln(&n)
	fmt.Printf("%d is prime or not prime  %t \n", n, isPrime(n))
}
