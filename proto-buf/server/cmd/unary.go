package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"server/gen/proto"
)

// User struct is a representation of user entity
type User struct {
	Name  string   `json:"name" validate:"required"`        // Name of the user, a required field
	Email string   `json:"email" validate:"required,email"` // Email ID of the user, should be in valid email format
	Roles []string `json:"roles" validate:"required"`       // Roles assigned to the user, a required field
}

// Signup method is part of the User Service defined in the proto file
// It receives a context and Signup Request and returns a Signup Response or an error
func (us *userService) Signup(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	nu := req.GetUser() // Fetching the user request sent by the client
	if nu == nil {
		return nil, status.Error(codes.Internal, "please provide required fields in correct format")

	}
	// Parsing the received request to our User struct
	u := User{
		Name:  nu.Name,
		Email: nu.Email,
		Roles: nu.Roles,
	}

	v := validator.New() // Creating a new validator instance

	// Validating the user struct according to the rules defined in the User struct.
	err := v.Struct(u)
	if err != nil {
		// If validation fails, return an error message with the error status.
		return nil, status.Error(codes.Internal, "please provide required fields in correct format")
	}

	// Assume you call your DB layer here and the data is stored in it.
	// Now you return back from there and start sending the response

	log.Println(u) // Logging the data for debugging purposes.

	// Returning the Signup Response with the result showing the email ID with account created message.
	return &proto.SignupResponse{Result: u.Email + " account created"}, nil
}
