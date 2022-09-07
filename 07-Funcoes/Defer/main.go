package main

import (
	"fmt"
	"io"
	"os"
)

/*
 A função defer faz com que uma determina função seja executada ao final no função
 em que ela esteja contida.
*/

func main() {
	ordemDeExecucao()
	fmt.Println("------------")

	execFuncoes()
	fmt.Println("------------")

	if err := escreverEmArquivo("exemplo.txt", "Go go go!!!"); err != nil {
		fmt.Println(err)
	}
}

// Os defer são alocados em uma pilha, de forma que o primeiro a ser declarado é o ultimo a ser executado.
func ordemDeExecucao() {
	defer fmt.Println("1º") // /\ ordem de execução
	defer fmt.Println("2º") // ||
	defer fmt.Println("3º") // ||
	defer fmt.Println("4º") // ||
	defer fmt.Println("5º") // ||

	fmt.Println("main")
}

func execFuncoes() {
	msg := "msg 1"

	// Ao passar pela declaração, armazenamos o valor atual de MSG dentro do escopo da função
	// pois ela é passada como paramentro
	defer func(m string) { fmt.Println("defer 1:", m) }(msg)

	// Já aqui, temos o ultimo favor que foi atribuido a MSG
	defer func() { fmt.Println("defer 2:", msg) }()
	msg = "msg 2"
}

// Podemo usa-ló para fechar recursos
func escreverEmArquivo(nomeDoArquivo string, msg string) error {
	arquivo, err := os.Create(nomeDoArquivo)
	if err != nil {
		return err
	}

	defer arquivo.Close()

	_, err = io.WriteString(arquivo, msg)
	if err != nil {
		return err
	}

	return arquivo.Close()
}
