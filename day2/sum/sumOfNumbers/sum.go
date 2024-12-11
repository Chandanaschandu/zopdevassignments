package sumOfNumbers

func Sum(n int) int {
	var Result int

	for i := 1; i <= n; i++ {
		Result = Result + i
	}

	return Result

}
