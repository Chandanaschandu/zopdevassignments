package employeeName

import (
	"testing"
	"time"
)

func TestGreetEmployee(t *testing.T) {
	employee := Employee{
		FirstName: "Chandana",
		LastName:  "S",
	}

	expected := "Hello Chandana S"
	result := GreetEmployee(employee)
	if result != expected {
		t.Errorf("GreetEmployee(%v) = %q; want %q", employee, result, expected)
	}

}
func BenchmarkGreetEmployee(b *testing.B) {
	employee := Employee{
		FirstName: "Chandana",
		LastName:  "S",
	}

	//expected := "Hello Chandana S"
	for range b.N {
		GreetEmployee(employee)
	}
}
func TestCalculateAge(t *testing.T) {

	dob := time.Date(2002, 11, 25, 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(2024, 12, 11, 0, 0, 0, 0, time.UTC)
	expectedAge := 22

	age, err := CalculateAge(dob, currentDate)
	if err != nil {
		t.Errorf("CalculateAge() returned an error: %v", err)
	}
	if age != expectedAge {
		t.Errorf("CalculateAge() = %d; want %d", age, expectedAge)
	}
}
func BenchmarkCalculateAge(b *testing.B) {
	dob := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(2024, 12, 11, 0, 0, 0, 0, time.UTC)

	// Run the benchmark loop.
	for range b.N {
		CalculateAge(dob, currentDate)
	}
}
