package main

import (
	"context"
	"grpc-test/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.ArithmeticServiceServer
}

func (s *server) Compute(ctx context.Context, req *pb.ArithmeticParamsRequest) (*pb.ArithmeticParamsResponse, error) {
	return &pb.ArithmeticParamsResponse{
		ParamResponse: req.Param1 + req.Param2,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
