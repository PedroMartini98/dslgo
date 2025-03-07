package main

import "errors"

/* Assignment
Implement the updateBalance function. It should take a customer pointer and a transaction, and return an error.
 Depending on the transactionType, it should either add or subtract the amount from the customers balance.
  If the customer does not have enough money, it should return the error insufficient funds. If the transactionType
  isn't a withdrawal or deposit, it should return the error unknown transaction type. Otherwise,
   it should process the transaction and return nil.

alice := customer{id: 1, balance: 100.0}
deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

updateBalance(&alice, deposit)
// id: 1 balance: 150 */

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line
func updateBalance(c *customer, t transaction) error {
	if t.transactionType == "deposit" {
		c.balance += t.amount
		return nil
	} else if t.transactionType == "withdrawal" {
		result := c.balance - t.amount
		if result < 0 {
			return errors.New("insufficient funds")
		}
		c.balance = result
		return nil
	} else {
		return errors.New("unknown transaction type")
	}
}

// ?
