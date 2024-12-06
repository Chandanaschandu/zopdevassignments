package main

import "fmt"

func SliceToMap(slice []int) map[int]int {

	result := make(map[int]int)

	for i, value := range slice {
		result[value] = i
	}

	return result
}

func main() {
	// Example usage
	s := []int{1, 2, 3, 4}

	result := SliceToMap(s)

	fmt.Println(result)
}
