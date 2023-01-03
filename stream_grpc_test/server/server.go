package main

import (
	"context"
	"fmt"
	"log"
	"microservice/stream_grpc_test/proto"
	"net"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)

const PORT = ":6868"

type server struct {
}

// 服务端流模式
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	var i int

	for {
		i++

		err := res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%d: %v", i, time.Now().Unix()),
		})

		if err != nil {
			log.Fatal("server send err: ", err)
		}

		time.Sleep(time.Second)
		if i >= 10 {
			break
		}
	}

	fmt.Println("server send finished")
	return nil
}

// 客户端流模式
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		resData, err := cliStr.Recv()
		if err != nil {
			fmt.Println("cliStr Recv err: ", err)
			break
		}

		fmt.Println(resData.Data)
	}

	fmt.Println("server receive finished")
	return nil
}

// 双向流模式
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			resData, err := allStr.Recv()
			if err != nil {
				fmt.Println("allStr Recv err: ", err)
				return
			}

			fmt.Println("收到客户端消息: ", resData.Data)
		}
	}()

	go func() {
		defer wg.Done()
		var i int

		for {
			i++

			var single []*proto.StreamResData_Result
			err := allStr.Send(&proto.StreamResData{
				Data: fmt.Sprintf("服务端发送的%d", i),
				Res:  single,
				Mp: map[string]string{
					"abc": "def",
				},
				AddTime: timestamppb.New(time.Now()),
			})
			if err != nil {
				fmt.Println("allStr Send err: ", err)
				return
			}
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("server finished")
	return nil
}

func (s *server) Ping(ctx context.Context, empty *empty.Empty) (*proto.Pong, error) {
	return &proto.Pong{}, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("listen err: ", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, new(server))
	err = s.Serve(listener)
	if err != nil {
		log.Fatal("grpc start err: ", err)
	}
}
