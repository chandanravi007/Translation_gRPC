package main

import (
	"context"
	"fmt"
	"log"
	"net"

	protos "github.com/chandanravi007/gRPC/UNARY/protos/protofiles"
	"google.golang.org/grpc"
)

type server struct {
	protos.UnimplementedGreetServiceServer
}

func (s server) Greet(ctx context.Context, req *protos.GreetRequest) (*protos.GreetResponse, error) {

	log.Println("User name: ", req.UserName)
	log.Println("Country code: ", req.CountryCode)

	var greeting string

	switch req.CountryCode {
	case "us":
		greeting = "hello " + req.UserName
	case "mx":
		greeting = "Hola " + req.UserName
	default:
		greeting = "Hola/Hello " + req.UserName
	}
	return &protos.GreetResponse{
		Result: greeting,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("starting server")

	s := grpc.NewServer()
	protos.RegisterGreetServiceServer(s, server{})

	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
