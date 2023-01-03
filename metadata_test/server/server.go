package main

import (
	"context"
	"fmt"
	"log"
	"microservice/grpc_test/proto"
	"net"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(c context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(c)
	if ok {
		for i, v := range md {
			fmt.Println(i, v)
		}
	}

	return &proto.HelloReply{
		Msg: "hello: " + request.Name,
	}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterGreeterServer(server, new(Server))

	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatal("listen err: ", err)
	}

	err = server.Serve(listen)
	if err != nil {
		log.Fatal("grpc serve err: ", err)
	}
}
