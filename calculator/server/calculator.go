package main

import (
	"context"
	"log"

	pb "github.com/watariRyo/gRPC-study/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNum + in.SecondNum,
	}, nil
}
