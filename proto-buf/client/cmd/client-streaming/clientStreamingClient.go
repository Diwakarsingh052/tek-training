package main

import (
	pb "client/gen/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

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
	// Create a context with a timeout of 20 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// Ensure that the cancel function is called no matter how the code exits
	defer cancel()

	// Call the CreatePost method of the client and get a reply stream, along with any error
	stream, err := client.CreatePost(ctx)
	// If there was an error in calling CreatePost method, then log and exit
	if err != nil {
		log.Fatalf("failed to call createPost server: %v", err)
	}

	// Simulate first batch of posts
	batch1 := []*pb.Post{
		{
			Title:  "The Science of Design",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "The Politics of Power",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "The Art of Programming",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}

	// Put the batch in a CreatePostRequest object
	p := &pb.CreatePostRequest{Posts: batch1}

	// Attempt to send CreatePost request through the stream
	err = stream.Send(p)
	// If there was an error in sending  request, then log and exit
	if err != nil {
		log.Fatalf("Failed to createPost request: %v", err)
	}

	// Adding latency, this simulates network delay or processing times
	time.Sleep(4 * time.Second)

	// Simulate second batch of posts
	batch2 := []*pb.Post{
		{
			Title:  "Post 11",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "Post 21",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "Post 31",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}

	// Put the second batch in a CreatePostRequest object
	p = &pb.CreatePostRequest{Posts: batch2}

	// Attempt to send CreatePost request through the stream
	err = stream.Send(p)
	// If there was an error in sending second batch request, then log and exit
	if err != nil {
		log.Fatalf("Failed to createPost request: %v", err)
	}

	// Close the client streaming and receive the server's response
	response, err := stream.CloseAndRecv()
	// If there was an error in receiving response, then log and exit
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}
	// Log the response from the server
	log.Printf("Response: %s", response.Result)
}
