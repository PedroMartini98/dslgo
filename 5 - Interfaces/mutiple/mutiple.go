package main

import "fmt"

/* Assignment
Complete the required methods so that the email type implements both the expense and formatter interfaces.

cost()
If the email is not "subscribed", then the cost is 5 cents for each character in the body.
If it is, then the cost is 2 cents per character.

Return the total cost of the entire email in cents.

format()
The format method should return a string in this format:

'CONTENT' | Subscribed

If the email is not subscribed, change the second part to "Not Subscribed":

'CONTENT' | Not Subscribed

The single quotes are included in the string, and CONTENT is the email's body. For example:

'Hello, World!' | Subscribed

Note: you may want to import the fmt package and use Sprintf. */

func (e email) cost() int {
	var cost int = 5
	if e.isSubscribed == true {
		cost = 2
	}
	var newCost int = len(e.body) * cost
	return newCost
}

func (e email) format() string {
	var isSub string = "Not Subscribed"
	if e.isSubscribed == true {
		isSub = "Subscribed"
	}

	return fmt.Sprintf("'%s' | %s", e.body, isSub)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
