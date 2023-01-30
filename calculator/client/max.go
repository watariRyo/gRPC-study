package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/watariRyo/gRPC-study/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatal("Error while calling Max: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 10},
		{Number: 5},
		{Number: 20},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving %v\n", err)
				break
			}

			log.Printf("Received %v\n", res)
		}

		close(waitc)
	}()

	<-waitc
}