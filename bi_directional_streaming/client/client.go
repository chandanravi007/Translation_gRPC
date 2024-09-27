package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/chandanravi007/gRPC/bi_directional_streaming/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error connecting to gRPC server: ", err.Error())
	}
	defer conn.Close()
	client := pb.NewChatserviceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}
	go func() {
		for {
			response, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}
			log.Printf("Server says: %s", response.Message)
		}
	}()
	//sending messages to server
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter messages (type 'exit' to quit):")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}
		err := stream.Send(&pb.ChatMessage{
			User:    "client",
			Message: text,
		})
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
	}
	stream.CloseSend()
	time.Sleep(time.Second)
}
