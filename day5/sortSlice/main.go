package main

import "fmt"

func sortSlice(ar []int, compare func(int, int) bool) {
	for i := 0; i < len(ar)-1; i++ {

		for j := i + 1; j < len(ar); j++ {
			if compare(ar[i], ar[j]) {
				ar[j], ar[i] = ar[i], ar[j]
			}
		}

	}
}

func compare(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}
func main() {

	ar := []int{2, 1, 9, 3, 5}
	sortSlice(ar, compare)
	fmt.Println(ar)
}
