package main

import (
	"google.golang.org/grpc"
	"log"
	."github.com/advanced-go-programming-book-code/ch4/s04/e03grpc-pubsub/helloservice"
	"context"
	"io"
	"fmt"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewPubsubServiceClient(conn)
	stream, err := client.SubscribeTopic(context.Background(), &String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}