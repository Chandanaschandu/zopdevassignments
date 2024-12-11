package primetest

func isPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}

	for i := 2; i*i <= n; i = i + 6 {
		if n%i == 0 || n%i+1 == 0 {
			return false
		}
	}

	return true
}
