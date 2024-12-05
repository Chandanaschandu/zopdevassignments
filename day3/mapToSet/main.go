package main

import (
	"fmt"
)

// create an empty set

func createSet() map[int]struct{} {
	return make(map[int]struct{})
}

// adding an element to the set

func addElement(set map[int]struct{}, element int) {
	set[element] = struct{}{}
}

// Delete an element from the set
func deleteElement(set map[int]struct{}, element int) {
	delete(set, element)
}

func main() {

	set := createSet()

	var n int
	fmt.Println("Enter the number of elements")
	fmt.Scanln(&n)

	var element int
	fmt.Println("Enter the elements in the set")
	for i := 0; i < n; i++ {
		fmt.Scanln(&element)
		addElement(set, element)
	}

	fmt.Println("Set after adding elements:", set)

	var del int
	fmt.Println("Number to be deleted")
	fmt.Scanln(&del)

	deleteElement(set, del)
	fmt.Printf("Set after deleting %d:%d", del, set)

}
