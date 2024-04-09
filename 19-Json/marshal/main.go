package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*
Marshall é o nome que se dar a conversão de uma struct ou map em Go para Json.
Fazemos então o "encode" dos dados de GO -> JSON
*/
func main() {
	p := Pessoa{"jp", 18, true}
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	// o valor de b está em bytes, então vamos cria uma visualização desses dados
	buffer := bytes.NewBuffer(b)
	fmt.Println(buffer.String())

	pessoas := map[int]Pessoa{
		1: p,
	}
	ps, err := json.Marshal(pessoas)
	if err != nil {
		panic(err)
	}

	// o valor de b está em bytes, então vamos cria uma visualização desses dados
	buffer = bytes.NewBuffer(ps)
	fmt.Println(buffer.String())
}

type Pessoa struct {
	Nome         string `json:"nome"`
	Idade        int    `json:"idade"`
	MaiorDeIdade bool   `json:"maiorDeIdade"`
}
