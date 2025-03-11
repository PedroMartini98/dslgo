package main

/* Assignment
Implement the processMessages function, it takes a slice of messages. It should process each message concurrently within a
 goroutine. Use the process function for this to simulate the benefit of using goroutines. Use a channel to ensure that
 all messages are processed and collected correctly, then return the slice of processed messages.
*/

import "time"

func processMessages(messages []string) []string {
	// ?
	ch := make(chan string, len(messages))
	for _, msg := range messages {
		go func(m string) {
			ch <- process(m)
		}(msg)
	}
	output := make([]string, 0, len(messages))
	for i := 0; i < len(messages); i++ {
		output = append(output, <-ch)
	}
	close(ch)
	return output
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
