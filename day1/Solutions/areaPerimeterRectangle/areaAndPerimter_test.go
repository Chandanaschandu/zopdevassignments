package areaPerimeterRectangle

import (
	"testing"
)

func TestAreaRectangle(t *testing.T) {
	tests := []struct {
		l      float64
		b      float64
		output float64
	}{
		{2, 3, 6.0},
		{1, 5, 5.0},
	}

	for _, test := range tests {
		res := AreaRectangle(test.l, test.b)
		if res != test.output {
			t.Errorf("Your result is %f and expected is %f", res, test.output)
		}
	}

}
func TestPeriRectangle(t *testing.T) {
	tests1 := []struct {
		a       float64
		b       float64
		output1 float64
	}{
		{2, 3, 10},
		{3, 4, 14},
	}
	for _, v := range tests1 {
		res1 := PeriRectangle(v.a, v.b)
		if res1 != v.output1 {
			t.Errorf("Your result is %f and expected is %f", res1, v.output1)
		}
	}
}
func BenchmarkAreaRectangle(b *testing.B) {
	tests := []struct {
		l      float64
		b      float64
		output float64
	}{
		{2, 3, 6.0},
		{1, 5, 5.0},
	}

	for range b.N {
		for _, test := range tests {
			AreaRectangle(test.l, test.b)
		}
	}

}
func BenchmarkPeriRectangle(b *testing.B) {
	tests1 := []struct {
		a       float64
		b       float64
		output1 float64
	}{
		{2, 3, 10},
		{3, 4, 14},
	}
	for range b.N {
		for _, test := range tests1 {
			PeriRectangle(test.a, test.b)
		}
	}
}
