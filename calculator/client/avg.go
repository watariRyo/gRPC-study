package main

import (
	"context"
	"log"

	pb "github.com/watariRyo/gRPC-study/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	numbers := []int32{1, 2, 3, 4, 5}

	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}