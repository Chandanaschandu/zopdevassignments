package employeeName

import (
	"fmt"
	"time"
)

type Employee struct {
	FirstName string
	LastName  string
	DOB       time.Time
}

func GreetEmployee(e Employee) string {
	return fmt.Sprintf("Hello %s %s", e.FirstName, e.LastName)
}
func CalculateAge(dob time.Time, currentDate time.Time) (int, error) {
	if dob.After(currentDate) {
		return 0, fmt.Errorf("birth date cannot be in the future")
	}

	age := currentDate.Year() - dob.Year()
	if currentDate.Month() < dob.Month() || (currentDate.Month() == dob.Month() && currentDate.Day() < dob.Day()) {
		age--
	}
	return age, nil
}
