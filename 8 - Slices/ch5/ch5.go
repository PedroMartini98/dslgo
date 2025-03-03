package main

import "strings"

/* Assignment
Complete the tagMessages function. It should take a slice of sms messages, and a function
(that takes a sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.
It should loop through each message and set the tags to the result of the passed in function.
Be sure to modify the messages of the original slice.
Ensure the tags field contains an initialized slice. No nil slices.
Complete the tagger function. It should take an sms message and return a slice of strings. */

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {

	for i := range messages {
		messages[i].tags = tagger(messages[i])
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	contentLower := strings.ToLower(msg.content)
	if strings.Contains(contentLower, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(contentLower, "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
