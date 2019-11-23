package main

import (
	"calculator/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Calculator Client")

	cc, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doSum(c)
	doClientStreaming(c)
	doBiStreaming(c)
}

func doBiStreaming(c calculatorpb.CalculatorServiceClient){
	fmt.Println("----------Bi Straming----------")

	stream, err := c.GetMax(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			fmt.Printf("[GET_MAX]Sending number: %v\n", number)
			_ = stream.Send(&calculatorpb.FindMaxRequest{
				Num: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		_ = stream.CloseSend()
	}()
	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Problem while reading server stream: %v", err)
				break
			}
			maximum := res.GetMaxNum()
			fmt.Printf("[GET_MAX]Received a new maximum of...: %v\n", maximum)
		}
		close(waitc)
	}()
	<-waitc
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("----------Straming CLient----------")

	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	numbers := []int32{3, 1, 8, 5, 9, 10, 34, 12, 99}
	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		_ = stream.Send(&calculatorpb.FindMaxRequest{
			Num: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[Find_MAX]The Max is: %v\n", res.GetMaxNum())
}


func doSum(c calculatorpb.CalculatorServiceClient)  {
	fmt.Println("----Sum---------")
	req := &calculatorpb.SumRequest{
		FirstNum: 100,
		SecondNum: 23,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling sum rpc: %v", err)
	}
	log.Printf("[SUM]Response from Sum: %v", res.SumResult)
}