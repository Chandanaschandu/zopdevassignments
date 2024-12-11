package calculatorMethods

import (
	"testing"
)

// TestAdd tests the add method of the Calculator struct.
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"Add positive numbers", 3, 2, 5},
		{"Add negative numbers", -3, -2, -5},
	}

	for _, tt := range tests {

		calc := Calculator{a: tt.a, b: tt.b}
		result := calc.add()
		if result != tt.expected {
			t.Errorf("add(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
		}

	}
}

// TestSub tests the sub method of the Calculator struct.
func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"Subtract positive numbers", 5, 3, 2},
		{"Subtract negative numbers", -5, -3, -2},
	}

	for _, tt := range tests {

		calc := Calculator{a: tt.a, b: tt.b}
		result := calc.sub()
		if result != tt.expected {
			t.Errorf("sub(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
		}

	}
}

// TestMul tests the mul method of the Calculator struct.
func TestMul(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"Multiply positive numbers", 2, 3, 6},
		{"Multiply negative numbers", -2, -3, 6},
	}

	for _, tt := range tests {
		{
			calc := Calculator{a: tt.a, b: tt.b}
			result := calc.mul()
			if result != tt.expected {
				t.Errorf("mul(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		}
	}
}

// TestDiv tests the div method of the Calculator struct.
func TestDiv(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		expected float64
	}{
		{"Divide positive numbers", 6, 3, 2},

		{"Divide by zero", 5, 0, 0}, // Expecting 0 due to the implementation handling division by zero
	}

	for _, tt := range tests {

		calc := Calculator{a: tt.a, b: tt.b}
		result := calc.div()
		if tt.b == 0 && result != 0 {
			t.Errorf("div(%f, %f) should return 0 when dividing by zero, got %f", tt.a, tt.b, result)
		} else if result != tt.expected {
			t.Errorf("div(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
		}
		
	}
}
