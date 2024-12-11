package Reverse

func ReverseSlice(s []int) []int {

	Start := 0
	Stop := len(s) - 1
	for Start < Stop {
		s[Start], s[Stop] = s[Stop], s[Start]
		Start++
		Stop--
	}

	return s

}
