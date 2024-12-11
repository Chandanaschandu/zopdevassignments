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
	// Create an employee instance for testing
	employee := Employee{
		FirstName: "John",
		LastName:  "Doe",
		DOB:       time.Date(2000, time.June, 15, 0, 0, 0, 0, time.UTC), // DOB for testing
	}

	expected := 24

	result := calculateAge(employee)

	if result != expected {
		t.Errorf("calculateAge(%v) = %d; want %d", employee, result, expected)
	}
}
func BenchmarkCalculateAge(b *testing.B) {

	employee := Employee{
		FirstName: "John",
		LastName:  "Doe",
		DOB:       time.Date(2000, time.June, 15, 0, 0, 0, 0, time.UTC),
	}
	for range b.N {
		calculateAge(employee)
	}
}
