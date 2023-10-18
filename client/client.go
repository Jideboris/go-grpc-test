package main

import (
	"context"
	"fmt"
	"grpc-test/pb"
	"log"
	"os"
	"strconv"

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
	params := os.Args[1:]
	fmt.Println("params", params)
	paramOne , _ := strconv.Atoi(params[0])
	paramTwo , _ := strconv.Atoi(params[1])

	request := &pb.ArithmeticParamsRequest{Param1: int32(paramOne), Param2: int32(paramTwo)}

	resp, err := client.Compute(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Receive response", resp.ParamResponse)
}
