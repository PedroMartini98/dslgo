package main

/* Assignment
Fix the bug in the code so that setMessage sets the message field of the given email structure,
and the new value persists outside the scope of the setMessage method. */

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}
