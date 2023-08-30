package controller

import (
	"regexp"
)

func validateName(name string) bool {
    if len(name) < 3 {
        return false
    }
    nameRegex := regexp.MustCompile(`^[A-Za-z\s]+$`)
    return nameRegex.MatchString(name)
}

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func validatePhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^\d{11}$`)
	return phoneRegex.MatchString(phone)
}