package models

import (
	"fmt"
	"regexp"
)

type User struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

func (u *User) Validate() error {
	emailRe := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	email := regexp.MustCompile(emailRe)
	isEmailValid := email.MatchString(u.Email)

	if !isEmailValid {
		return fmt.Errorf("%T: %s", isEmailValid, u.Email)
	}
	phoneRe := `^\+?[0-9]{10,15}$`
	phone := regexp.MustCompile(phoneRe)
	isPhoneValid := phone.MatchString(u.PhoneNumber)

	if !isPhoneValid {
		return fmt.Errorf("%T:%s", isPhoneValid, u.PhoneNumber)
	}

	return nil
}
