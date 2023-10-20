package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
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

	req := &pb.GetPostsRequest{UserId: 101}
	stream, err := client.GetPosts(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		//receiving values from stream
		post, err := stream.Recv()

		//if the server has finished sending the request, we will quit
		if err == io.EOF {
			break
		}
		//any other kind of error would be caught here
		if err != nil {
			log.Println(err)
			return
		}
		select {
		case <-stream.Context().Done():
			fmt.Println("server cancelled")
		default:

		}
		fmt.Println("reading stream")
		//printing data received
		fmt.Println(post)
		fmt.Println()

	}
}
