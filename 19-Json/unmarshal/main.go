package main

import (
	"encoding/json"
	"fmt"
)

/*
Unmarshall é o nome que se dar a conversão de um JSON ou map ou struct em Go.
Fazemos então o "decode" dos dados de JSON -> GO
*/
func main() {
	pessoaJSON := `{"nome":"jp","idade":18,"maiorDeIdade":true}`

	var pessoa Pessoa
	err := json.Unmarshal([]byte(pessoaJSON), &pessoa) // o segundo paramentro deve ser um ponteiro pata que os dados sejam setados nele
	if err != nil {
		panic(err)
	}

	fmt.Println(pessoa)

	pessoasJSON := `{"1":{"nome":"jp","idade":18,"maiorDeIdade":true}}`
	var pessoas map[int]Pessoa

	err = json.Unmarshal([]byte(pessoasJSON), &pessoas)
	if err != nil {
		panic(err)
	}

	fmt.Println(pessoas)
}

type Pessoa struct {
	Nome         string `json:"nome"`
	Idade        int    `json:"idade"`
	MaiorDeIdade bool   `json:"maiorDeIdade"`
}
