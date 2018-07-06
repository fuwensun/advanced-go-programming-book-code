package main

import (
	"net/rpc"
	"log"
	"fmt"
)

func main() {
	//rpc 连线，生成*Client
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	//Client 远程调用HelloService.Hello方法，参数"hello"
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}