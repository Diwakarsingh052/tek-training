package main

import (
	"fmt"
	pb "proto-buf-basics/proto-buf-basics/proto"
)

func main() {
	enum()
}
func simpleMessage() {

	r := pb.BlogRequest{
		BlogId:  101,
		Title:   "Introduction to Protocol Buffers",
		Content: "Test",
	}
	fmt.Println(r.GetBlogId(), r.GetContent())
	//return a string representation of a Protobuf message.
	//This is useful for debugging and logging purposes
	fmt.Println(r.String())

	//reset the values of a Protobuf message to their default values.
	//This is useful when you want to reuse a message without having to create a new one from scratch
	r.Reset()
	fmt.Println(r.String())

}

func enum() {
	p := pb.Product{
		ProductId: "101",
		Category:  pb.Category_CATEGORY_CLOTHING, // only one category could be set
	}

	//r:= proto.Restaurant{Order: proto.}
	fmt.Println(p.GetCategory())
}
