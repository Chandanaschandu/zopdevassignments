package main

import (
	"fmt"
	"github.com/Chandanaschandu/training/day2"
)

func main() {

	//Day2
	var n int
	fmt.Println("enter the number")
	// Day 1
	fmt.Scanln(&n)
	fmt.Printf("%d is prime number %t \n", n, day2.IsPrime(n)) //To check number is prime or not

	//calculator
	fmt.Println("enter two numbers a and b")
	var a, b int

	fmt.Scanln(&a, &b)

	fmt.Printf("The calculator number is %d\n", day2.Calculator(a, b))
	// Day2
	fmt.Println("Enter the number to sum")
	var m int
	fmt.Scanln(&m)

	fmt.Printf("Sum of the given number is %d", day2.Sum(m)) //to calculate the sum of the number

}
