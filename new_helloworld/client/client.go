package main

import (
	"fmt"
	"log"
	"microservice/new_helloworld/client_proxy"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "127.0.0.1:1234")

	var reply string
	err := client.Hello("momozi", &reply)
	if err != nil {
		log.Fatal("call err: ", err)
	}

	fmt.Println(reply)
}
