package main

import (
	"context"
	"fmt"
	"log"
	"microservice/grpc_test/proto"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

type customCredential struct {
}

func (c *customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "123456",
		"appkey": "qwertyuiop",
	}, nil
}

func (c *customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	//interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	//	start := time.Now()
	//
	//	md := metadata.New(map[string]string{
	//		"appid":  "123456",
	//		"appkey": "qwertyuiop",
	//	})
	//	ctx = metadata.NewOutgoingContext(ctx, md)
	//
	//	err = invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Printf("本次请求耗时: %s\n", time.Since(start))
	//	return
	//}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	opts = append(opts, grpc.WithPerRPCCredentials(&customCredential{}))
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
