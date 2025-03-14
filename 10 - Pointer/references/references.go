package main

/* Assignment
Complete the removeProfanity function.

It should use the strings.
ReplaceAll function to replace all instances of the following words in the input message with asterisks.

Word	Replacement
fubb	****
shiz	****
witch	*****
It should update the value in the pointer and return nothing.

Do not alter the function signature. */

import (
	"strings"
)

func removeProfanity(message *string) {
	output := *message
	output = strings.ReplaceAll(output, "fubb", "****")
	output = strings.ReplaceAll(output, "shiz", "****")
	output = strings.ReplaceAll(output, "witch", "*****")
	*message = output
}
