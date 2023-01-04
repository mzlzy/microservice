package main

import (
	"context"
	"fmt"
	"log"
	"microservice/grpc_test/proto"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		start := time.Now()
		err = invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("本次请求耗时: %s\n", time.Since(start))
		return
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial("127.0.0.1:8989", opts...)
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
