package primetest

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		des    string
		input  int
		output bool
	}{
		{"Is prime", 3, true},
		{"Not prime", 4, false},
	}
	for _, test := range tests {
		res := isPrime(test.input)
		if res != test.output {
			t.Errorf("Yours output is %t and expected is %t", res, test.output)
		}
	}
}
func BenchmarkIsPrime(b *testing.B) {
	tests := []struct {
		des    string
		input  int
		output bool
	}{
		{"Is prime", 3, true},
		{"Not prime", 4, false},
	}
	for range b.N {
		for _, test := range tests {
			isPrime(test.input)
		}
	}
}
