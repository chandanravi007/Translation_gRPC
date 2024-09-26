package main

import (
	"context"
	"log"
	"time"

	pb "github.com/chandanravi007/gRPC/clientstreaming/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("0.0.0.0:50057", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error connecting to gRPC server: ", err.Error())
	}
	defer conn.Close()

	client := pb.NewPhoneClient(conn)
	ClientStreamNumCheck(client)
}
func ClientStreamNumCheck(client pb.PhoneClient) {
	requests := []*pb.NumCheckRequest{
		{
			Number: "30",
		},
		{
			Number: "23",
		},
	}
	respStream, err := client.NumCheck(context.Background())
	if err != nil {
		log.Fatalf("error while calling clinet %v", err)
	}
	for _, req := range requests {
		log.Printf("sending req %v\n", req)
		err := respStream.Send(req)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	res, err := respStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving the response %v\n", err)
	}
	log.Println(res.CheckResult)
}
