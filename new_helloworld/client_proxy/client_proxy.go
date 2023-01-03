package client_proxy

import (
	"log"
	"microservice/new_helloworld/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	dial, err := rpc.Dial(protocol, address)
	if err != nil {
		log.Fatal("get client err: ", err)
	}

	return HelloServiceStub{dial}
}

func (s *HelloServiceStub) Hello(request string, reply *string) error {
	return s.Call(handler.HelloServiceName+".Hello", request, reply)
}
