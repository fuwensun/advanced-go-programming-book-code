package main

import (
	"net/rpc"
	"log"
	"fmt"
)
const HelloServiceName = "path/to/pkg.HelloService"

func main() {
	//rpc 连线，生成*Client
	client, err := rpc.Dial("tcp", "localhost:1234")	//v1
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	//Client 远程调用HelloService.Hello方法，输入参数"hello"，输出参数&reply
	//err = client.Call("HelloService.Hello", "hello", &reply)				//v1
	err = client.Call(HelloServiceName+".Hello", "hello", &reply)	//v2
	if err != nil {
		log.Fatal(err)
	}

	//打印回复
	fmt.Println(reply)
}