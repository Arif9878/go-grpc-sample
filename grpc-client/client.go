package main

import (
	"context"
	"log"

	"github.com/Arif9878/go-grpc-sample/grpc-server/chat"
	"google.golang.org/grpc"
)

const (
	address = ":9000"
)

func main() {
	// Set up a connection to the server.
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
