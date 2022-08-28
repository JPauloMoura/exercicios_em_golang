package main

import (
	"fmt"
	"reflect"
)

func variaveis() {
	var inteiro int32
	inteiro = 10
	fmt.Println("Inteiro: ", inteiro)

	var nome, sobrenome string = "jose", "Silva"
	fmt.Printf("nome:%s da %s \n", nome, sobrenome)

	sexo := 'M'
	fmt.Printf("sexo: %v, type: %T \n", sexo, sexo)

	resp := true
	fmt.Println("resp: ", resp, "type:", reflect.TypeOf(resp))
}

func capturaInput() {

	fmt.Print("Digite seu nome: ")
	var nome string
	fmt.Scan(&nome) // &nome -> [endereço]

	fmt.Print("Digite sua idade: ")
	var idade int
	fmt.Scan(&idade) // &idade -> [endereço]

	if idade == 0 {
		fmt.Printf("Olá %s, você digitou uma idade inválida", nome)
	} else {
		fmt.Printf("Olá %s, eu também tenho %d anos", nome, idade)
	}
}

func main() {
	capturaInput()
}
