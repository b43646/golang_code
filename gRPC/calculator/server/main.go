package main

import (
	"calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {}

func (server) GetMax(stream calculatorpb.CalculatorService_GetMaxServer) error {
	fmt.Println("------------GETMAX-----------")
	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {

			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		num := req.GetNum()

		if num > maximum {
			maximum = num
			sendErr := stream.Send(&calculatorpb.FindMaxResponse{
				MaxNum:maximum,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}

func (server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	ret := req.FirstNum + req.SecondNum
	return &calculatorpb.SumResponse{SumResult:ret}, nil
}

func (server) FindMax(stream calculatorpb.CalculatorService_FindMaxServer) error {
	fmt.Println("Received FindMaximum RPC")
	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			sendErr := stream.SendMsg(&calculatorpb.FindMaxResponse{
				MaxNum:maximum,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		num := req.GetNum()

		if num > maximum {
			maximum = num
		}
	}
}

func main() {
	fmt.Println("Calculator Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
