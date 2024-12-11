package CountNumberofwords

import (
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"Good morning", 2},
		{"Hello how are you", 4},
	}

	for _, test := range tests {
		result := CountWords(test.input)
		if result != test.output {
			t.Errorf("CountWords(%q) = %d; want %d", test.input, result, test.output)
		}
	}

}
func BenchmarkCountWords(b *testing.B) {
	tests := []struct {
		input  string
		output int
	}{
		{"Good morning", 2},
		{"Hello how are you", 4},
	}
	for range b.N {
		for _, test := range tests {
			CountWords(test.input)
		}
	}

}
