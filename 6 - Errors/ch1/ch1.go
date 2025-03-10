package main

import "errors"

/* Complete the validateStatus function. It should return an error when the status update violates any of the rules:

If the status is empty, return an error that reads status cannot be empty.
If the status exceeds 140 characters, return an error that says status exceeds 140 characters. */

func validateStatus(status string) error {
	length := len(status)

	if length == 0 {
		return errors.New("status cannot be empty")
	}

	if length > 140 {
		return errors.New("status exceed 140 characters")
	}

	return nil
}
