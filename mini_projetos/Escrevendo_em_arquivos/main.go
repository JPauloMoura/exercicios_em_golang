package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const permisao = 0755
const dataAtual = "02/01/2006" //seguindo o padrão Go
const hhAtual = "15:04:05"     //seguindo o padrão Go
func main() {
	logs := []string{"log 1", "log 2", "log 3"}

	for _, v := range logs {
		salvarLogs(v)
	}

	lerArquivo()

}

func salvarLogs(log string) {
	arquivo, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, permisao)
	if err != nil {
		fmt.Println("err->", err)
		return
	}

	horario := time.Now().Format("[" + dataAtual + " " + hhAtual + "] ")
	arquivo.WriteString(horario + log + "\n")
}

func lerArquivo() {
	conteudo, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("err->", err)
		return
	}

	fmt.Println(string(conteudo))
}
