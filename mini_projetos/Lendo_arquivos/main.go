package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// lerArquivo1()
	// lerArquivo2()
	test()
}

type File struct {
	Name  string
	bytes *os.File
}

func start() *File {
	ponteiro, _ := os.Open("file.txt")

	return &File{
		Name:  "file 1",
		bytes: ponteiro,
	}
}

func test() {

	// file := start()
	// defer file.bytes.Close()

	b, err := ioutil.ReadFile("./file.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
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
