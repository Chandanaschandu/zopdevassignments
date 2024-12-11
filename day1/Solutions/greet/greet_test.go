package greet

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		desc   string
		input  string
		output string
	}{
		{"Greet  the person", "Good morning", "Good morning"},
		{"Greet  the person", "Good evening", "Good evening"},
		{"Greet the person", "Hello", "Hi"},
	}
	for _, test := range tests {
		res := Greet(test.input)
		if res != test.output {
			t.Errorf("The %s is %s and output is %s", test.desc, test.input, test.output)
		}
	}

}
func BenchmarkGreet(b *testing.B) {
	tests := []struct {
		desc   string
		input  string
		output string
	}{
		{"Greet of the person", "Good morning", "Good morning"},
		{"Greet of the person", "Good evening", "Good evening"},
	}
	for range b.N {
		for _, test := range tests {
			Greet(test.input)
		}
	}
}
