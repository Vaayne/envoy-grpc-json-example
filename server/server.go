package main

import (
	"context"
	"log"

	"server/pb/bookstore"
	"server/pb/helloworld"
)

// helloworldServer is used to implement helloworld.GreeterServer.
type helloworldServer struct {
	helloworld.UnimplementedGreeterServer
}

func (s *helloworldServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	if err := in.ValidateAll(); err != nil {
		log.Printf("validate error: %v", err)
		return nil, err
	}
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *helloworldServer) Status(ctx context.Context, in *helloworld.Empty) (*helloworld.StatusResponse, error) {
	return &helloworld.StatusResponse{Status: "ok "}, nil
}

type bookstoreServer struct {
	bookstore.UnimplementedBookStoreServer
}

func (s *bookstoreServer) Status(ctx context.Context, in *bookstore.Empty) (*bookstore.StatusResponse, error) {
	return &bookstore.StatusResponse{Status: "bookstore ok"}, nil
}
