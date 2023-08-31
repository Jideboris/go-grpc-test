package main

import (
	"context"
	"fmt"
	"grpc-test/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Anthony Bborisade Client is Ready ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewArithmeticServiceClient(cc)
	request := &pb.ArithmeticParamsRequest{Param1: 2,Param2:6}

	resp, err := client.Compute(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Receive response", resp.ParamResponse)
}
