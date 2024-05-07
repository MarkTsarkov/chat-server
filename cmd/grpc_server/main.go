package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/marktsarkov/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()

	reflection.Register(s)

	desc.RegisterChatServer(s, &server{})

	log.Printf("Server listening at %v", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var port = "50052"

type server struct {
	desc.UnimplementedChatServer
}


func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (resp *desc.CreateResponse, err error) {
	log.Printf("Users: %v", req.Username)

	return &desc.CreateResponse{
		Id: 0,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleted chat id: %v", req.Id)

	return nil, nil
}

func (s *server) Send(ctx context.Context, req *desc.SendRequest) (resp *desc.SendResponse, err error) {
	log.Printf("Message send by: %v", req.From)
	return nil, nil
}


