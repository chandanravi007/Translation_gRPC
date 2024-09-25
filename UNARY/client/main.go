package main

import (
	"context"
	"log"

	protos "github.com/chandanravi007/Translation_gRPC/UNARY/protos/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer cc.Close()

	c := protos.NewGreetServiceClient(cc)

	getGreeting("Jack", "us", c)
	getGreeting("Jose", "mx", c)

}

func getGreeting(name, countryCode string, c protos.GreetServiceClient) {

	log.Println("creating greeting")

	res, err := c.Greet(context.Background(), &protos.GreetRequest{
		CountryCode: countryCode,
		UserName:    name,
	})

	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}
	log.Println(res.Result)
}
