## Grpc
É um framework (conjunto de ferramentas) desenvolvido pela Google para facilitar
o processo de comunicação entre sistemas. 

- Indicado para comunicação entre microsseviços por ser leve (trafega binarios ao invés de texto);
- Utiliza o protocolo HTT2, onde um mesma canal envia e recebe multiplas request e response;

A comunicação grpc usa um protocolo de comunicação chamado Protol Buffer ou Protobuf.
Esse é um contrado usado para serializa e desserializar os dados binários que são enviados e recebidos entre as partes.

Os contratos são chamados de protofile e tem mais ou menos esse formato:
```go
synta = "proto3"
message SearchRequest{
    string query = 1;
    int32 page = 2;
    int32 result_per_page = 3;
}
``````

1) Site da documentação e guia de iniciação do gRpc https://grpc.io/docs/languages/go/quickstart/
    Aqui você também vai encontra o tutorial de instalação
    
2) Site da documentação e guia de iniciação do protocal buffer https://protobuf.dev/programming-guides/proto3/

## Geração dos arquivos
Dado um arquivo .proto com as messagens e services que iremos usar, geramos os stubs com o seguinte comando:
`$protoc --go_out=. --go-grpc_out=. hello.proto`
