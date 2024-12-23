package models

import (
	"fmt"
	"regexp"
)

type Users struct {
	UserName    string `json:"user_name"`
	UserAge     int    `json:"user_age"`
	Phonenumber string `json:"phone_Number"`
	Email       string `json:"email"`
}

func (u *Users) Validate() error {
	emailRe := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	email := regexp.MustCompile(emailRe)
	isEmailValid := email.MatchString(u.Email)

	if !isEmailValid {
		return fmt.Errorf("Invalid email id")
	}

	phoneRe := `^\+?[0-9]{10,15}$`
	phone := regexp.MustCompile(phoneRe)
	isPhoneValid := phone.MatchString(u.Phonenumber)

	if !isPhoneValid {
		return fmt.Errorf("Invalid Phone number")
	}

	return nil

}
