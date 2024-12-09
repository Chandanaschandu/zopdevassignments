package main

import "fmt"

var stack []int

func push(value int) {

	stack = append(stack, value)
}

func pop() (int, string) {
	if len(stack) < 0 {
		return -1, "Empty stack"
	}

	popped := len(stack) - 1
	stack = stack[:len(stack)-1]
	return popped, "nil"
}
func main() {

	push(20)
	push(40)
	push(30)

	fmt.Println("Stack after pushing elements", stack)

	pop()
	
	fmt.Println("Stack after popping elements", stack)

}
