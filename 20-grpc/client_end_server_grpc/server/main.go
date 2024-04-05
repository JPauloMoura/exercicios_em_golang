package main

import (
	"context"
	"exercicios/20-grpc/client_end_server_grpc/server/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3001")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterHandlerUserServer(s, &HandlerUser{})
	reflectionRefisterServer(s)

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

/*
Esta é uma instrução especial que expõe todos os serviços gRPC acessíveis ao público em um servidor gRPC.
Isso significa que qualquer pessoa pode solicitar ao seu servidor gRPC que emita detalhes dos métodos de serviço RPC.
Isso é usado em lugares onde você deseja chamar dinamicamente APIs gRPC.
Por dinamicamente quero dizer que o cliente não precisa manter as estruturas de dados ".proto" e registrar o stub do cliente RPC,
como em chamadas usando o "grpcurl".

	$ grpcurl -plaintext -d '{"name": "jp"}' localhost:3001 HandlerUser.Create
*/
func reflectionRefisterServer(s *grpc.Server) {
	reflection.Register(s)
}

var countUser int32

type HandlerUser struct {
	pb.UnimplementedHandlerUserServer
}

func (HandlerUser) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	countUser++
	return &pb.CreateUserResponse{Id: countUser, Name: req.Name}, nil
}
