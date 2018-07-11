package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"net/http"
	"io"
)

//服务接口定义
const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

//服务类型定义
type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}


func main() {
	//rpc服务注册
	rpc.RegisterName("HelloService", new(HelloService))

	//http注册handler 和pattern
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	//http服务添加监听端口
	http.ListenAndServe(":1234", nil)
}