package main

import (
	"net/rpc"
	"net"
	"log"
)

//服务接口定义
const HelloServiceName = "path/to/pkg.HelloService"

/*

type HelloServiceInterface = interface {				//v6
	Hello(request *String, reply *String) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
*/

type HelloService struct{}


func (p *HelloService) Hello(request *String, reply *String) error {	//v6
	reply.Value = "hello:" + request.GetValue()
	return nil
}


func main() {

	//rpc服务注册
	//rpc.RegisterName("HelloService", new(HelloService)) 	//v1
	//RegisterHelloService(new(HelloService))					//v2
	var rpcs rpc.Server
	RegisterHelloService(&rpcs,new(HelloService))				//v8

	//*TCPListener 建立
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	//v4
	for {
		//*TCPConn  获取
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		//rpc服务绑定
		//go rpc.ServeConn(conn)
		go rpcs.ServeConn(conn)
	}
}

