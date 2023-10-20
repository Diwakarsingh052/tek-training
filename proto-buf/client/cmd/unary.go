package main

import (
	pb "client/gen/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

// the main function
func main() {
	// dialOptions for the gRPC connection
	dialOpts := []grpc.DialOption{
		// WithTransportCredentials specifies the transport credentials for the connection
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Dialing a connection to the gRPC server at localhost:5001
	conn, err := grpc.Dial("localhost:5001", dialOpts...)
	// Handle any connection error
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	// Ensure to close the connection when the main function finishes
	defer conn.Close()

	// Instantiate a client for the UserService service
	client := pb.NewUserServiceClient(conn)

	// Create a SignupRequest object
	req := &pb.SignupRequest{
		User: &pb.User{
			// Set the user name to John
			Name: "John",
			// Set the user email
			Email: "johnemail.com",
			// Set the password for the user
			Password: "abc",
			// Set the user roles to ADMIN and USER
			Roles: []string{"ADMIN", "USER"},
		},
	}

	// Create a context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// Cancel the context when the main function finishes
	defer cancel()

	// Execute the Signup request and get a response
	res, err := client.Signup(ctx, req)

	// Handle any errors that occurred during the Signup request
	if err != nil {
		// Try to extract a gRPC status from the error
		grpcStatus, ok := status.FromError(err)
		if ok {
			// Log the gRPC status code and message
			log.Println(grpcStatus.Code())
			log.Println(grpcStatus.Message())
			return
		}

		// Fail the program with a fatal error
		log.Fatalln(err)
	}

	// Log the result of the Signup request
	log.Println(res.Result)
}
