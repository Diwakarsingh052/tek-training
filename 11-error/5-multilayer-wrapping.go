package main

import (
	"errors"
	"fmt"
	"log"
)

// Declare global error values for different kinds of errors we might come across.

var ErrAdmission = errors.New("admission error")
var ErrFees = errors.New("fees error")
var ErrFinance = errors.New("finance not accepted")

func main() {
	err := admission()
	// If admission returns an error, it's logged here. This is where we first unwrap the error.
	if err != nil {
		log.Println(err)
	}
}

// Admission function simulates an admission process, and it calls the fees() function.
func admission() error {
	err := fees()
	// If there's an error in fees(), we wrap it here with ErrAdmission, adding another layer to our 'error chain'.
	if err != nil {
		return fmt.Errorf("%w %w", err, ErrAdmission)
	}
	// If no errors occurred in fees(), we return nil.
	return nil
}

// Fees function simulates a fee payment process, and it calls the finance() function.
func fees() error {
	err := finance()
	// If finance() returns an error, we wrap it here with ErrFees, starting the process of multi-layered error wrapping.
	if err != nil {
		return fmt.Errorf("%w %w", err, ErrFees)
	}
	// If no error occurred in finance(), we return nil.
	return nil
}

// Finance function simulates a financial operation. In this case it always returns an error.
func finance() error {
	// Here we're creating the first layer of our 'error chain' by wrapping the base error ErrFinance.
	return fmt.Errorf("%w", ErrFinance)
}
