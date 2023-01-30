package main

import (
	"io"
	"log"

	pb "github.com/watariRyo/gRPC-study/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client strseam: %v\n", err)
		}

		log.Printf("Receiving %v\n", req)
		sum += req.Number
		count++
	}
}