package sumOfNumbers

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		description string
		input       int
		output      int
	}{
		{"sum of 4 numbers", 4, 10},
		{"sum of 5 numbers", 5, 14},
	}
	for _, test := range tests {
		res := Sum(test.input)
		if res != test.output {
			t.Errorf("Your output is %d and expected output is %d", res, test.output)
		}
	}

}
func BenchmarkSum(b *testing.B) {
	tests := []struct {
		description string
		input       int
		output      int
	}{
		{"sum of 4 numbers", 4, 10},
		{"sum of 5 numbers", 5, 15},
	}
	for range b.N {
		for _, test := range tests {
			Sum(test.input)
		}
	}

}
