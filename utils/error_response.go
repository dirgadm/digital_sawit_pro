package utils

import (
	"errors"
	"regexp"
)

var ErrPhoneNumberAlreadyExists = errors.New("Phone number already exists")
var LengthPhoneNumber = "Length of phone number should be greater than 9 or less than 14 "
var LengthFullName = "Length of full name should be greater than 3 or less than 61"
var StartingCode = "Phone Number should start with +62"
var NotStrongPassword = "Password is not strong."

func IsStrongPassword(password string) bool {
	// Minimum length of 6 characters
	if len(password) < 6 {
		return false
	}

	// Maximum length of 64 characters
	if len(password) > 64 {
		return false
	}

	// At least 1 uppercase letter
	uppercaseRegex := regexp.MustCompile("[A-Z]")
	if !uppercaseRegex.MatchString(password) {
		return false
	}

	// At least 1 number
	numberRegex := regexp.MustCompile("[0-9]")
	if !numberRegex.MatchString(password) {
		return false
	}

	// At least 1 special (non-alphanumeric) character
	specialCharRegex := regexp.MustCompile("[!@#$%^&*()-_+=<>?]")
	if !specialCharRegex.MatchString(password) {
		return false
	}

	return true
}
