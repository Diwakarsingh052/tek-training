package middlewares

import (
	"errors"
	"service-app/auth"
)

type Mid struct {
	a *auth.Auth
}

func NewMid(a *auth.Auth) (Mid, error) {
	if a == nil {
		return Mid{}, errors.New("auth can't be nil")
	}
	return Mid{a: a}, nil
}
