package Reverse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		des    string
		input  []int
		output []int
	}{
		{"Given slice is", []int{1, 2, 3}, []int{3, 2, 1}},
		{"input slice", []int{9, 8, 7}, []int{7, 8, 9}},
	}

	for _, test := range tests {
		res := ReverseSlice(test.input)
		assert.Equal(t, res, test.output, "error")

	}

}
func BenchmarkReverseSlice(b *testing.B) {
	tests := []struct {
		des    string
		input  []int
		output []int
	}{
		{"Given slice is", []int{1, 2, 3}, []int{3, 2, 1}},
		{"input slice", []int{9, 8, 7}, []int{7, 8, 9}},
	}
	for range b.N {
		for _, test := range tests {
			ReverseSlice(test.input)
		}
	}
}
