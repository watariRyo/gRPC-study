package main

import (
	"context"
	"io"
	"log"

	pb "github.com/watariRyo/gRPC-study/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatal("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v¥n", err)
		}
		log.Printf("GreetManyTimes %d\n", msg.Result)
	}
}