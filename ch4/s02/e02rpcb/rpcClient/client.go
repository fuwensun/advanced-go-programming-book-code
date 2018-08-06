package main

import (
	"net/rpc"
	"log"
	"fmt"
	."github.com/advanced-go-programming-book-code/ch4/s02/e02rpcb/helloservice"
)
const HelloServiceName = "path/to/pkg.HelloService"


//客户类型定义
type HelloServiceClient struct {
	*rpc.Client
}
//var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)//v3
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

/*
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	//return p.Client.Call(HelloServiceName+".Hello", request, reply)

	var qq String
	var pp String
	qq.Value = request;
	err := p.Client.Call(HelloServiceName+".Hello", qq, &pp)
	if(err == nil){
		*reply = pp.GetValue()
	}
	return err
}
*/

func (p *HelloServiceClient) Hello(request String, reply *String) error {	//v6
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	//rpc 连线，生成*Client
	//client, err := rpc.Dial("tcp", "localhost:1234") //v1
	client, err := DialHelloService("tcp", "localhost:1234")//v3
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//var reply string
	//Client 远程调用HelloService.Hello方法，输入参数"hello"，输出参数&reply
	//err = client.Call("HelloService.Hello", "hello", &reply)			//v1
	//err = client.Call(HelloServiceName+".Hello", "hello", &reply)		//v2
	//err = client.Hello("hello", &reply)								//v3

	var request String
	var reply String
	request.Value = "hello";
	err = client.Hello(request, &reply)
	if err != nil {
		log.Fatal(err)
	}

	//打印回复
	fmt.Println(reply.GetValue())
}