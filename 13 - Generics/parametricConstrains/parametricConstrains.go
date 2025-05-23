package main

/* Assignment
The chief architect at Mailio has decided she wants to implement billing with generics. Specifically,
she wants us to create a new biller interface. A biller is an interface that can be used to charge a customer,
and it can also report its name.

There are two kinds of billers:

userBiller (cheaper)
orgBiller (more expensive)
A customer is either a user or an org. A user will be billed with a userBiller and an org with an orgBiller.

Create the new biller interface. It should have 2 methods:

Charge
Name
The good news is that the architect already wrote the userBiller and orgBiller types for us that fulfill
this new biller interface. Use the definitions of those types and their methods to figure out how to write the biller
interface definition. */

import (
	"fmt"
)

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

// don't edit below this line

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}
