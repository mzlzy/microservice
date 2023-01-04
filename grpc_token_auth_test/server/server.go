package main

import (
	"context"
	"fmt"
	"log"
	"microservice/grpc_test/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(c context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Msg: "hello: " + request.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了请求")
		resp, err = handler(ctx, req)
		fmt.Println("请求处理完毕")
		return
	}
	opt := grpc.UnaryInterceptor(interceptor)
	server := grpc.NewServer(opt)
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
