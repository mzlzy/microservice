package main

import (
	"log"
	"microservice/new_helloworld/handler"
	"microservice/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen err: ", err)
	}

	err = server_proxy.RegisterHelloService(&handler.HelloService{})
	if err != nil {
		log.Fatal("register server err: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept err: ", err)
		}

		go rpc.ServeConn(conn)
	}

}
