// Package models contains structures and functions for user data manipulation.
package models

import "errors"

// users is a map used as an in-memory database.
// The keys are user IDs and the values are User objects.
var users = map[uint64]User{
	123: {
		FName:    "Bob",
		LName:    "abc",
		Password: "someSecretPassword",
		Email:    "bob@email.com",
	},
}

// ErrDataNotPresent is an error returned when the requested user data is not found.
var ErrDataNotPresent = errors.New("data is not their")

// FetchUser retrieves a User object from the in-memory database by its ID.
// It returns the User object and an error, which is nil if the user is found and ErrDataNotPresent otherwise.
func FetchUser(userId uint64) (User, error) {
	value, ok := users[userId]

	// If ok is false, there is no user with the given ID.
	if !ok {
		return User{}, ErrDataNotPresent
	}

	// If ok is true, the user was found and is returned.
	return value, nil
}
