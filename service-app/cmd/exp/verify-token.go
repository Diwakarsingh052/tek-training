package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5" // Importing the JWT library for Go to handle JSON Web Tokens
	"log"                          // Log package for logging error messages
	"os"                           // OS package used for interacting with the OS such as reading files
)

// JWT token given as string
var tkn = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.
eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY5NzYyMzYyNCwiaWF0IjoxNjk3NjIwNjI0fQ.a6jFV8eBU6GhRya1tvdG5IOEs4D8tTkK1lvCkl1YIUBSV9QC8yMdy5F8pKZzzn1zrnJNy_EH-Wc9A8waByAv_TmLIDwXrhSgilCGyodbip-FhrjIvBb_n_GQG8JGC8Fvsdgue5Fg5ovgV6ApN795jX-RvNEdaHDyb3yAP1mMmtBMb7vnB4xACKG_h7J2KpDuEshBvB5k270rcJbD9TUK7R9G39tQVi-cx7N14OCVrFoUiRlhT8UM2UlrS2bIUpmyDMNR3H0WqeV-tKH-6M2KV-z0cbHsdwjywLwBpezshsM3xhIoTkKZDAvfDoWrM0uJB5C69YlLZtt6AqKJDOXkqQ`

func main() {
	// Reads the public key from pubkey.pem file
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		// If there's an error reading the file, print an error message and stop execution
		log.Fatalln("not able to read pem file")
	}

	// Parse the read public key to RSA public key format
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	if err != nil {
		// If there's an error parsing the public key, log the error and stop execution
		log.Fatalln(err)
	}
	var c jwt.RegisteredClaims
	// Parsing the JWT token with the claims
	token, err := jwt.ParseWithClaims(tkn, &c, func(token *jwt.Token) (interface{}, error) {
		// Provides the public key for validating the JWT token
		return publicKey, nil
	})

	if err != nil {
		// If error while parsing the token, print the error and exit
		log.Println("parsing token", err)
		return

	}
	if !token.Valid {
		// If the token is not valid, log the error and exit
		log.Println("invalid token")
		return
	}

	// Print the claims from the token
	fmt.Printf("%+v", c)

}
