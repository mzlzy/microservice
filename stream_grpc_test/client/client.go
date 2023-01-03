package main

import (
	"context"
	"fmt"
	"log"
	"microservice/stream_grpc_test/proto"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:6868", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("client dial err: ", err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)

	pong, _ := client.Ping(context.Background(), &empty.Empty{})
	fmt.Println(pong)

	// 服务端流模式
	res, err := client.GetStream(context.Background(), &proto.StreamReqData{Data: "client"})
	if err != nil {
		log.Fatal("client GetStream err: ", err)
	}
	for {
		resData, err := res.Recv()
		if err != nil {
			fmt.Println("client Recv err: ", err)
			break
		}

		fmt.Println(resData.Data)
	}

	// 客户端流模式
	putRes, err := client.PutStream(context.Background())
	if err != nil {
		log.Fatal("client PutStream err: ", err)
	}
	var i int
	for {
		i++
		err = putRes.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("this is send %d", i),
		})

		time.Sleep(time.Second)
		if i >= 10 {
			break
		}
	}

	// 客户端双向流
	allRes, err := client.AllStream(context.Background())
	if err != nil {
		log.Fatal("client AllStream err: ", err)
	}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			resData, err := allRes.Recv()
			if err != nil {
				fmt.Println("allRes Recv err: ", err)
				return
			}

			fmt.Println("收到服务端消息: ", resData.Data)
		}
	}()

	go func() {
		defer wg.Done()
		var i int

		for {
			i++
			err := allRes.Send(&proto.StreamReqData{Data: fmt.Sprintf("客户端发送的%d", i)})
			if err != nil {
				fmt.Println("allRes Send err: ", err)
				return
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("client finished")
}
