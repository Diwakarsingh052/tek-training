package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type ctxKey int

const Key ctxKey = 1

// Auth is a type that deals with authentication-related activities. It contains two fields, privateKey and publicKey which are
// used for token generation and verification respectively.
type Auth struct {
	privateKey *rsa.PrivateKey // privateKey is used to sign the JWT token.
	publicKey  *rsa.PublicKey  // publicKey is used to validate the JWT token.
}

// NewAuth is a constructor function for Auth struct. It accepts privateKey and publicKey as parameters and returns
// an instance of Auth struct. If either of privateKey or publicKey is nil, it returns an error.
func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Auth, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("private/public key cannot be nil")
	}
	return &Auth{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// GenerateToken is a method for Auth struct. It generates a new JWT token using the provided claims and
// signs it using the privateKey of the Auth struct it's called upon. If there is an error during signing,
// it returns an error.
func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {
	//NewWithClaims creates a new Token with the specified signing method and claims.
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Signing our token with our private key.
	tokenStr, err := tkn.SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}

	return tokenStr, nil
}

// ValidateToken is a method for Auth struct. It verifies the provided JWT token using the publicKey of the Auth struct
// it's called upon and returns the parsed claims if the JWT token is valid. If the JWT token is invalid or
// there is an error during parsing, it returns an error.
func (a *Auth) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	var c jwt.RegisteredClaims
	// Parse the token with the registered claims.
	tkn, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})
	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("parsing token %w", err)
	}
	// Check if the parsed token is valid.
	if !tkn.Valid {
		return jwt.RegisteredClaims{}, errors.New("invalid token")
	}
	return c, nil
}
