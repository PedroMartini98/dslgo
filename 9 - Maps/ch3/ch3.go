package main

/* Assignment
There is a bug in the logAndDelete function, fix it!

This function should always delete the user from the user's map,
 which is a map that stores the user's name as keys.
 It also returns a log string that indicates to the caller some information about the user's deletion.

To avoid bugs like this in the future, instead of calling delete before each return,
 just defer the delete once at the beginning of the function. */

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	user, ok := users[name]
	defer delete(users, name)
	if user.admin {
		return logAdmin
	}
	if !ok {
		return logNotFound
	}
	return logDeleted
}
