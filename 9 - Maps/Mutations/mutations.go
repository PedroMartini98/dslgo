package main

import "errors"

/* Assignment
It's important to keep up with privacy regulations and to respect our user's data.
We need a function that will delete user records.

Complete the deleteIfNecessary function.
The user struct has a scheduledForDeletion field that determines if they are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.*/

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user, ok := users[name]
	if ok == false {
		return false, errors.New("not found")
	}
	if user.scheduledForDeletion == false {
		return false, nil
	}
	delete(users, name)
	return true, nil

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
