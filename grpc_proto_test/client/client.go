package main

import (
	"context"
	"fmt"
	"log"
	"microservice/grpc_test/proto"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8989", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc dial err: ", err)
	}

	client := proto.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "momozi"})
	if err != nil {
		log.Fatal("SayHello err: ", err)
	}

	fmt.Println(reply.Msg)
}
