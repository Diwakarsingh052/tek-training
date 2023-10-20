package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"sync"
	"time"
)

// main function
func main() {
	// dial the gRPC server
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	// Close the connection before the main function exits
	defer conn.Close()

	// Use the connection to initialize a new UserService client
	client := pb.NewUserServiceClient(conn)

	// Context with a timeout of 20 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	// Cancel the context when the main function exits
	defer cancel()

	// Initiate the GreetEveryone stream
	stream, err := client.GreetEveryone(ctx)
	if err != nil {
		log.Fatalf("failed to call GreetEveryone stream: %v\n", err)
	}

	// List of requests to be sent
	requests := []*pb.GreetEveryoneRequest{
		{FirstName: "John"},
		{FirstName: "Bruce"},
		{FirstName: "Roy"},
	}

	// WaitGroup to synchronize goroutines
	wg := new(sync.WaitGroup)

	// Add count 2 for each of the two goroutines spawned
	wg.Add(2)

	// Goroutine for sending requests
	go func() {
		defer wg.Done()
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)
			err = stream.Send(req)
			if err != nil {
				log.Println(err)
				return
			}
			select {
			case <-stream.Context().Done():
				fmt.Println("server cancelled")
				return
			default:
				// do nothing if the context has not been cancelled
			}
		}

		// close the send direction of the stream
		err := stream.CloseSend()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	// Goroutine for receiving responses
	go func() {
		defer wg.Done()

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Printf("stream has ended") // received the end of the stream signal
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			select {
			case <-stream.Context().Done():
				fmt.Println("cancelled") // received cancellation signal
				return
			default:
				// do nothing if the context has not been cancelled
			}
			log.Printf("Received: %v\n", res.Result)
		}
	}()

	// Block until all goroutines have finished executing
	wg.Wait()
}
