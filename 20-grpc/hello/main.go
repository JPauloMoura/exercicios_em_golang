package main

import (
	"context"
	"exercicios/20-grpc/hello/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type handler struct {
	pb.UnimplementedHandlerHelloServer
}

func (h handler) Hello(ctx context.Context, user *pb.HelloUser) (*pb.Blank, error) {

	fmt.Println(user.UserName)
	return &pb.Blank{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
