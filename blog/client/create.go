package main

import (
	"context"
	"log"

	pb "github.com/watariRyo/gRPC-study/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("----createblog was invoked----")

	blog := &pb.Blog{
		AuthorId: "Bob",
		Title: "My First Blog",
		Content: "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("expected ")
	}

	log.Printf("Blog has benn created: %s\n", res)
	return res.Id
}