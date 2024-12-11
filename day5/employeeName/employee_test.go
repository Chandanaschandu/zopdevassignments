package employeeName

import "testing"

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
