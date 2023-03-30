package main

import (
	"context"
	"log"

	pb "github.com/nikzayn/blog-grpc/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("-----updateBlog was invoked------")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Pooja",
		Title:    "I am new here",
		Content:  "I will soon post my content here",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Printf("Blog was updated: %v\n", newBlog)
}
