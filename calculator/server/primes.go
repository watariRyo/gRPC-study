package main

import (
	"log"

	pb "github.com/watariRyo/gRPC-study/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, strean pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number % divisor == 0 {
			strean.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor ++
		}
	}

	return nil
}
