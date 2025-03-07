package main

/* Assignment
Let's make our profanity checker safe. Update the removeProfanity function. If message is nil,
 return early to avoid a panic. After all, there are no bad words to remove. */

import (
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}
