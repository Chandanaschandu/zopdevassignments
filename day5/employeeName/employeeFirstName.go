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
func calculateAge(emp Employee) int {
	currentYear := time.Now().Year()
	age := currentYear - emp.DOB.Year()

	// Adjust if the employee hasn't had their birthday yet this year
	if time.Now().YearDay() < emp.DOB.YearDay() {
		age--
	}

	return age
}
