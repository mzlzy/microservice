package main

import (
	"context"
	"fmt"
	"log"
	"microservice/grpc_test/proto"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8989", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc dial err: ", err)
	}

	client := proto.NewGreeterClient(conn)

	//md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02 15:04:05"), "timestamp", time.Now().Add(time.Hour).Format("2006-01-02 15:04:05"))
	md := metadata.New(map[string]string{
		"timestamp":         time.Now().Format("2006-01-02 15:04:05"),
		"timestamp_another": time.Now().Add(time.Hour).Format("2006-01-02 15:04:05"),
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	reply, err := client.SayHello(ctx, &proto.HelloRequest{Name: "momozi"})
	if err != nil {
		log.Fatal("SayHello err: ", err)
	}

	fmt.Println(reply.Msg)
}
