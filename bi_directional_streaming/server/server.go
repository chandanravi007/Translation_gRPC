package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/chandanravi007/gRPC/bi_directional_streaming/protofiles"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatserviceServer
}

func (s *server) Chat(stream pb.Chatservice_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("unable to serve the client at this momemnt %v", err)
		}
		log.Printf("Received message from %s: %s", msg.User, msg.Message)
		response := &pb.ChatMessage{
			User:    "server",
			Message: fmt.Sprintf("hello %s, you said: %s", msg.User, msg.Message),
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}
func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatserviceServer(s, &server{})
	log.Printf("Server is listening on port 50051...")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
