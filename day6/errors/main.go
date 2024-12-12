package main

import (
	"fmt"
	"time"
)

type Myerror struct {
	today time.Time
	day   string
}

func (t *Myerror) Error() string {
	return fmt.Sprintf("today time is %v and day is %s", t.today, t.day)
}

func run() error {
	return &Myerror{
		time.Now(),
		"Thursday",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
