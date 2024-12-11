package calculatorMethods

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	cal := Calculator{}
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{2, 3, 5},
		{7, 9, 16},
	}
	for _, test := range tests {
		res := cal.Add(test.a, test.b)
		if res != test.expected {
			t.Errorf("Your output is %f and expected output is %f", res, test.expected)

		}
	}

}
func TestCalculator_Sub(t *testing.T) {
	cal := Calculator{}
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{5, 3, 2},
	}
	for _, test := range tests {
		res := cal.Sub(test.a, test.b)
		if res != test.expected {
			t.Errorf("Your output is %f and expected output is %f", res, test.expected)

		}
	}
}
func TestCalculator_Mul(t *testing.T) {
	cal := Calculator{}
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{5, 2, 10},
	}
	for _, test := range tests {
		res := cal.Mul(test.a, test.b)
		if res != test.expected {
			t.Errorf("Your output is %f and expected output is %f", res, test.expected)

		}
	}
}
func TestCalculator_Div(t *testing.T) {
	cal := Calculator{}
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{5, 2, 2.5},
	}
	for _, test := range tests {
		res := cal.Div(test.a, test.b)
		if res != test.expected {
			t.Errorf("Your output is %f and expected output is %f", res, test.expected)

		}
	}
}
func BenchmarkCalculator_Add(b *testing.B) {
	cal := Calculator{}
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{2, 3, 5},
		{7, 9, 16},
	}
	for range b.N {
		for _, test := range tests {
			cal.Add(test.a, test.b)
		}
	}
}
func BenchmarkCalculator_Sub(b *testing.B) {
	cal := Calculator{}
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{2, 3, -1},
		{9, 5, 4},
	}
	for range b.N {
		for _, test := range tests {
			cal.Sub(test.a, test.b)
		}
	}
}
func BenchmarkCalculator_Mul(b *testing.B) {
	cal := Calculator{}
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{2, 3, 6},
		{9, 5, 45},
	}
	for range b.N {
		for _, test := range tests {
			cal.Mul(test.a, test.b)
		}
	}
}
func BenchmarkCalculator_Div(b *testing.B) {
	cal := Calculator{}
	tests := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{5, 2, 2.5},
		{4, 0, 0},
	}
	for range b.N {
		for _, test := range tests {
			cal.Div(test.a, test.b)
		}
	}
}
