package main

import (
	"log"
	"net"

	pb "github.com/chandanravi007/gRPC/ServerStreaming/protofiles"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStreamingServiceServer
}

func (s server) GetDataStreaming(request *pb.DataRequest, srv pb.StreamingService_GetDataStreamingServer) error {
	log.Println("Fetch data streaming")
	for i := 0; i < 10; i++ {
		value := rangStringBytes(500)
		resp := pb.DataResponse{
			Part:   int32(i),
			Buffer: value,
		}
		if err := srv.Send(&resp); err != nil {
			log.Println("error generating response")
			return err
		}
	}
	return nil
}
func rangStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	// create listener
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic("error building server: " + err.Error())
	}

	// create gRPC server
	s := grpc.NewServer()
	pb.RegisterStreamingServiceServer(s, server{})

	log.Println("start server")

	if err := s.Serve(listener); err != nil {
		panic("error building server: " + err.Error())
	}
}
