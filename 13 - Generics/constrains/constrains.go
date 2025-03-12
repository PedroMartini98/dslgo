package main

/* Assignment
We have different kinds of "line items" that we charge our customer's credit cards for.
Line items can be things like "subscriptions" or "one-time payments" for email usage.


If they do have enough funds:

Add the line item to the user's history by appending the newItem to the slice of oldItems.
 This new slice is your first return value.
Calculate the user's new balance by subtracting the cost of the new item from their balance.
This is your second return value.*/

import (
	"errors"
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	var zero []T

	if newItem.GetCost()-balance > 0 {
		return zero, 0.0, errors.New("insufficient funds")
	} else {
		balance -= newItem.GetCost()
		oldItems = append(oldItems, newItem)
		return oldItems, balance, nil
	}
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}
