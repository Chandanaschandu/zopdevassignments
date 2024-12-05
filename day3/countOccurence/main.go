package main

import "fmt"

func countCharacters(word string) map[string]int {
	frequency := make(map[string]int)

	for _, char := range word {
		frequency[string(char)] = frequency[string(char)] + 1
	}

	return frequency
}

func main() {
	var word string

	fmt.Scanln(&word)

	fmt.Println("Frequency of Word is : ", countCharacters(word))
}
