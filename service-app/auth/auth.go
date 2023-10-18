package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type ctxKey int

const Key ctxKey = 1

// Auth struct privateKey field would be used to verify and generate token
type Auth struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Auth, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("private/public key cannot be nil")
	}
	return &Auth{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {
	//NewWithClaims creates a new Token with the specified signing method and claims
	//NewWithClaims creates a new Token with the specified signing method and claims
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	//signing our token with our private key
	tokenStr, err := tkn.SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}

	return tokenStr, nil

}

func (a *Auth) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	var c jwt.RegisteredClaims
	tkn, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})
	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("parsing token %w", err)
	}
	if !tkn.Valid {
		return jwt.RegisteredClaims{}, errors.New("invalid token")
	}
	return c, nil

}
