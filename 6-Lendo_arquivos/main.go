package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// lerArquivo1()
	lerArquivo2()
}

func lerArquivo1() {
	ponteiro, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	//mostrando conteudo até o caracter '4'
	arquivo := bufio.NewReader(ponteiro)
	conteudo, err := arquivo.ReadString('4') //ReadString recebe o local do texto para finalizar a leitura
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(conteudo))

	ponteiro.Close()
}

func lerArquivo2() {
	ponteiro, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	arquivo := bufio.NewReader(ponteiro)

	for {
		linha, err := arquivo.ReadString('\n')
		linha = strings.TrimSpace(linha)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(linha)
		if err == io.EOF {
			fmt.Println(linha)
			break
		}

	}
}
