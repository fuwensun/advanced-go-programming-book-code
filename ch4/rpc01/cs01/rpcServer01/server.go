package main

import (
	"net/rpc"
	"net"
	"log"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}


func main() {

	//rpc 注册
	rpc.RegisterName("HelloService", new(HelloService))

	//*TCPListener 建立
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	//*TCPConn  获取
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	
	//rpc绑定
	rpc.ServeConn(conn)
}