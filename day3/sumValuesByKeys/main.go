package main

import "fmt"

func sumValues(m map[string][]int) map[string]int {
	result := make(map[string]int)

	for key, values := range m {
		var sum int
		for _, value := range values {
			sum = sum + value
		}
		result[key] = sum
	}

	return result

}

func main() {
	m := map[string][]int{
		"a": {1, 2, 3},
		"b": {8, 3, 4},
		"c": {1, 3, 4},
	}

	fmt.Println("The Sum of the values:", sumValues(m))
}
