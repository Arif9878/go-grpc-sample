package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Arif9878/go-grpc-sample/grpc-server/chat"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	chat.UnimplementedChatServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &chat.Message{Body: "Hello From the Server!"}, nil
}

const (
	port = ":9000"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	chat.RegisterChatServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
