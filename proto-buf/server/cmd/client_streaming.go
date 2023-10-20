package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	pb "server/gen/proto"
	"sync"
)

// In client-streaming RPC, the client sends multiple messages/request to the server
// instead of a single request.
// The server sends back a single response to the client.

func (us *userService) CreatePost(stream pb.UserService_CreatePostServer) error {

	wg := &sync.WaitGroup{}
	// Receive CreatePost request from client in batches
	for {
		//receiving the request from the client
		req, err := stream.Recv()

		//If the client has finished sending the request, we will quit
		if err == io.EOF {
			log.Println("stream ended")
			break
		}
		if err != nil {
			log.Println(" error in stream")
			return err
		}

		//latency of server processing
		//in the meantime when server was doing the processing of the request ,
		//if the request is cancelled then we would detect that using select
		//time.Sleep(time.Second * 4)
		//during the request if client close the connection we will know inside this select block
		select {
		case <-stream.Context().Done():
			log.Println("client cancelled the request")
			return errors.New("client disconnected")
		default:
			// Client is still connected
		}

		// Process create post request
		b, _ := json.MarshalIndent(req, "", " ")
		log.Printf("Received Create Post Requests: %v", string(b))

		posts := req.GetPosts()
		log.Println("adding all the posts into the db")
		wg.Add(1)

		//add posts in db
		go AddPost(posts, wg)
	}

	wg.Wait()
	// Return response
	return stream.SendAndClose(&pb.CreatePostResponse{Result: "done"})

}

func AddPost(posts []*pb.Post, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, p := range posts {

		log.Println("adding post ", p.Title)
	}
}
