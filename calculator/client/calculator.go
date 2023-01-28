package main

import (
	"context"
	"log"

	pb "github.com/watariRyo/gRPC-study/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNum: 10,
		SecondNum: 20,
	})

	if err != nil {
		log.Fatal("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}