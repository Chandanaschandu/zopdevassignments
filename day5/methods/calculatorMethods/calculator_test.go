package calculatorMethods

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{3, 2, 5.0},
		{-1, -1, -2.0},
	}

	for _, tt := range tests {
		c := Calculator{a: tt.a, b: tt.b}
		result := c.add()
		if result != tt.expected {
			t.Errorf("Expected %.2f, but got %.2f", tt.expected, result)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{5, 3, 2.0},
		{10, 20, -10.0},
		{-5, -3, -2.0},
	}

	for _, tt := range tests {
		c := Calculator{a: tt.a, b: tt.b}
		result := c.sub()
		if result != tt.expected {
			t.Errorf("Expected %.2f, but got %.2f", tt.expected, result)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{4, 3, 12.0},
		{0, 5, 0.0},
		{-4, 3, -12.0},
	}

	for _, tt := range tests {

		c := Calculator{a: tt.a, b: tt.b}
		result := c.mul()
		if result != tt.expected {
			t.Errorf("Expected %.2f, but got %.2f", tt.expected, result)
		}
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		a, b       float64
		expected   float64
		shouldFail bool
	}{
		{6, 2, 3.0, false},
		{6, 0, 0.0, true},
	}

	for _, tt := range tests {

		c := Calculator{a: tt.a, b: tt.b}

		var result float64
		if tt.b == 0 {

			if tt.shouldFail {

				result = 0.0
			}
		} else {

			result = c.div()
		}

		if tt.shouldFail && tt.b != 0 {
			t.Errorf("Expected division by zero, but the division was successful")
		}

		if !tt.shouldFail && result != tt.expected {
			t.Errorf("Expected %.2f, but got %.2f", tt.expected, result)
		}

	}
}

func BenchmarkAdd(b *testing.B) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{3, 2, 5.0},
		{-1, -1, -2.0},
	}

	for range b.N {
		for _, tt := range tests {
			c := Calculator{a: tt.a, b: tt.b}
			c.add()

		}
	}
}

func BenchmarkSub(b *testing.B) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{5, 3, 2.0},
		{10, 20, -10.0},
		{-5, -3, -2.0},
	}
	for range b.N {
		for _, tt := range tests {
			c := Calculator{a: tt.a, b: tt.b}
			c.sub()
		}

	}
}

func BenchmarkMul(b *testing.B) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{4, 3, 12.0},
		{0, 5, 0.0},
		{-4, 3, -12.0},
	}
	for range b.N {
		for _, tt := range tests {
			c := Calculator{a: tt.a, b: tt.b}
			c.mul()
		}

	}
}

func BenchmarkDiv(b *testing.B) {
	tests := []struct {
		a, b       float64
		expected   float64
		shouldFail bool
	}{
		{6, 2, 3.0, false},
		{6, 0, 0.0, true},
	}

	for range b.N {
		for _, tt := range tests {

			c := Calculator{a: tt.a, b: tt.b}
			var _ float64
			if tt.b == 0 {

				if tt.shouldFail {

					_ = 0.0
				}
			} else {

				c.div()
			}
		}
	}
}
