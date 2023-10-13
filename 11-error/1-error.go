// Main package to group our functions
package main

// Importing required packages
import (
	"errors" // Package for manipulating errors
	"fmt"
	"log"
)

// A map named 'user' holding int as key and string as value
var user = make(map[int]string)

// ErrRecordNotFound Custom error message for record not found error
var ErrRecordNotFound = errors.New("record not found")

// FetchRecord Function tries to retrieve a value from the 'user' map
// If record is not found, it returns an error
// If record is found, it returns the retrieved value and a nil error
func FetchRecord(id int) (string, error) {

	name, ok := user[id] // Retrieving user record
	if !ok {             // If record isn't found
		return "", ErrRecordNotFound // Return custom error
	}
	return name, nil // Record found, return it
}

func main() {

	// Fetching record with ID 10
	n, err := FetchRecord(10)
	if err != nil { // If error occurred while fetching record
		log.Println(err) // Log the error and return
		return
	}
	// Record found, print the name
	fmt.Println(n)
}
