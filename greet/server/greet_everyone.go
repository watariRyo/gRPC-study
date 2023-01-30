package main

import (
	"io"
	"log"

	pb "github.com/watariRyo/gRPC-study/greet/proto"
)

func (s *Server)GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello" + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Fatalf("Error while sending data %v\n", err)
		}
		log.Printf("GreetEveryone: %v\n", res)
	}
}