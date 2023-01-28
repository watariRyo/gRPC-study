package main

import (
	"context"
	"log"

	pb "github.com/watariRyo/gRPC-study/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Bob",
	})

	if err != nil {
		log.Fatal("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}