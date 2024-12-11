package CountNumberofwords

import "strings"

func CountWords(sentence string) int {
	// Split the sentence into words by spaces and filter empty strings.

	return len(strings.Fields(sentence))
}
