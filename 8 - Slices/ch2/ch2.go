package main

/* Assignment
Implement the isValidPassword function by looping through each character in the password string.
 Make sure the password is long enough and includes at least one uppercase letter and one digit.

Assume passwords consist of ASCII characters only. */

import "unicode"

func isValidPassword(password string) bool {
	length := len(password)
	if length > 12 {
		return false
	}
	if length < 5 {
		return false
	}

	hasUpper := false
	hasDigit := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsDigit(char) {
			hasDigit = true
		}
		// Early exit if both are found
		if hasUpper && hasDigit {
			return true
		}
	}
	if hasDigit && hasUpper {
		return true
	}
	return false
}
