package main

import (
	"time"
	"google.golang.org/grpc"
	"net"
	"log"
	."github.com/advanced-go-programming-book-code/ch4/s04/e03grpc-pubsub/helloservice"
	"github.com/docker/docker/pkg/pubsub"
	"context"
	"strings"
	"fmt"
)


type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(
	ctx context.Context, arg *String,
) (*String, error) {
	p.pub.Publish(arg.GetValue())
	//return &String{}, nil

	//debug
	reply := &String{Value: "<Publish>  " + arg.GetValue()}
	fmt.Println(reply.GetValue())
	return reply, nil
	//debug
}


func (p *PubsubService) Subscribe(
	arg *String, stream PubsubService_SubscribeServer,
) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			//debug
			fmt.Printf("<debug> %t %s %s %t\n",
				ok,arg.GetValue(),key,strings.HasPrefix(key,arg.GetValue()))
			//debug
			if strings.HasPrefix(key,arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterPubsubServiceServer(grpcServer,NewPubsubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)

}
