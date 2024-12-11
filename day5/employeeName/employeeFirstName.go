package employeeName

import (
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
}

func GreetEmployee(e Employee) string {
	return fmt.Sprintf("Hello %s %s", e.FirstName, e.LastName)
}
