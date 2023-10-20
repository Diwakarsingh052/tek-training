package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"server/gen/proto"
)

// userService struct that contains an unimplemented UserServiceServer defined in proto file
type userService struct {
	proto.UnimplementedUserServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Println(err)
		return
	}

	//NewServer creates a gRPC server which has no service registered
	//and has not started to accept requests yet.
	s := grpc.NewServer()
	//registers a gRPC service implementation with a gRPC server.
	//The second argument is a pointer to the service implementation struct,
	//which implements the methods defined in the gRPC service interface.
	proto.RegisterUserServiceServer(s, &userService{})

	//exposing gRPC service to be tested by postman
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		fmt.Println(err)
		return
	}
}
