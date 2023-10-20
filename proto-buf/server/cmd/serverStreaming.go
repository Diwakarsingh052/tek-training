package main

import (
	"log"
	pb "server/gen/proto"
)

func (us *userService) GetPosts(req *pb.GetPostsRequest, stream pb.UserService_GetPostsServer) error {
	id := req.GetUserId()
	log.Println("GetPosts: ", "fetching all posts for user id ", id)

	//assume these posts we are getting in batches
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

	err := stream.Send(&pb.GetPostsResponse{Posts: batch1})
	if err != nil {
		log.Println(err)
		return err
	}

	//constructing the second batch
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

	b2 := &pb.GetPostsResponse{Posts: batch2}
	//sending first batch of response
	err = stream.Send(b2)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("all posts are sent for user id", id)
	return nil

}
