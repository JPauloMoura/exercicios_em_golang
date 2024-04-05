package main

import (
	"context"
	"exercicios/20-grpc/client_end_server_grpc/server/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Quer dizer que nossa conexão não utilizará um certificado de segurança
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	baseUrl := "localhost:3001"

	conn, err := grpc.NewClient(baseUrl, credentials)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// NewHandlerUserClient devolve uma instancia com todos os metodos do serviço que padem ser usados
	handlerUser := pb.NewHandlerUserClient(conn)
	body := &pb.CreateUserRequest{
		Name: "fulano",
	}

	resp, err := handlerUser.Create(context.Background(), body)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
