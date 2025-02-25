package main

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {
	charLimit := 100
	if membershipType == "premium" {
		charLimit = 1000
	}

	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: charLimit,
		},
	}

}

type Membership struct {
	Type             string
	MessageCharLimit int
}

/*
Assignment
Create a new struct called Membership, it should have:

A Type string field
A MessageCharLimit integer field
Update the User struct to embed a Membership.

Complete the newUser function. It should return a new User with all the fields set as you would expect
based on the inputs. If the user is a "premium" member,
the MessageCharLimit should be 1000, otherwise, it should only be 100.*/
