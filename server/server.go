package main

import (
	"context"
	proto "go-grpc/proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct{}

func main() {
	listner, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listner); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetA()
	b := request.GetB()
	c := a + b
	return &proto.Response{Result: c}, nil

}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetA()
	b := request.GetB()
	c := a * b
	return &proto.Response{Result: c}, nil

}
