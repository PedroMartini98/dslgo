package main

/* Assignment
Your task is to implement a function that filters a slice of messages based on the message type.

Complete the filterMessages function.
It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link").
It should return a new slice of Message interfaces containing only messages of the specified type. */

func filterMessages(message []Message, typeOfMessage string) []Message {
	output := make([]Message, 0)
	for i, m := range message {
		if m.Type() == typeOfMessage {
			output = append(output, message[i])
		}
	}
	return output
}

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender      string
	MediaType   string
	Description string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender string
	URL    string
	Title  string
}

func (lm LinkMessage) Type() string {
	return "link"
}
