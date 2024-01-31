package main

import (
	"log"
	"net"

	"github.com/usmonzodasomon/grpc-service/pkg/adder"
	proto "github.com/usmonzodasomon/grpc-service/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	srv := &adder.GRPCServer{}

	proto.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
