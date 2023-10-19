package middlewares

import (
	"errors"
	"service-app/auth"
)

// Mid is a structure that holds an authenticated session.
// This is typically used for maintaining user sessions or secure transactions.
type Mid struct {
	// 'a' attribute is a pointer to an 'Auth' object.
	// It's important to note that 'a'
	//is a pointer because we want to refer to the original 'Auth' object and not a COPY of it.
	a *auth.Auth
}

// NewMid is a function which takes an 'Auth' object pointer
// and returns a Mid instance and an error.
// Purpose of this function is to initialize
// and return a new instance of 'Mid' structure.
func NewMid(a *auth.Auth) (Mid, error) {
	// It first checks if 'a' is nil
	// 'a' should not be nil because 'nil' indicates that the 'Auth' object does not exist.
	if a == nil {
		// An error is returned when 'a' is 'nil'.
		return Mid{}, errors.New("auth can't be nil")
	}
	//If 'a' is not 'nil', a new 'Mid' instance is returned with 'a' as a field.
	// A nil error is returned, indicating that there were no issues with the initialization.
	return Mid{a: a}, nil
}
