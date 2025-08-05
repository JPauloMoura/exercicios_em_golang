package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Vamos aprender sonre a manipução de arquivos em Go.

O arquivo é criado no no diretorio em que o programa foi chamado.
*/
func main() {
	filename := "teste.pdf"
	file, err := criarArquivo(filename)
	if err != nil {
		panic(err)
	}

	err = escreverEmUmArquivo(file, []byte("PDF com o conteúdo utilizado!"))
	if err != nil {
		panic(err)
	}

	file.Close()

	err = printConteudoDoArquivo(filename)
	if err != nil {
		panic(err)
	}

	err = leituraDeGrandesArquivos(filename)
	if err != nil {
		panic(err)
	}

	err = deletarArquivo(filename)
	if err != nil {
		panic(err)
	}
}

func criarArquivo(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar arquivo %s: %s", filename, err.Error())
	}

	return file, nil
}

func escreverEmUmArquivo(file *os.File, content []byte) error {
	_, err := file.Write(content)
	if err != nil {
		return fmt.Errorf("falha ao escrever no arquivo %s: %s", file.Name(), err.Error())
	}

	return nil
}

func printConteudoDoArquivo(filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("falha ao ler arquivo %s: %s", filename, err.Error())
	}
	fmt.Println(string(b))
	return nil
}

/*
Existem situações em que temos que lidar com a leitura de arquivos bem grandes, mas nossa aplicação não possuem RAM suficiente.
Nesse caso temos que ler pedaços do arquivo, e fazemos isso utilizando buffers;
*/
func leituraDeGrandesArquivos(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("falha ao abrir arquivo %s: %s", filename, err.Error())
	}

	defer file.Close()

	reader := bufio.NewReader(file) // cria um "leitor" para lidar com o arquivo
	buffer := make([]byte, 3)       // cria um buffer com 10bytes, que é o tamanho das partes que seram lidas

	for {
		position, err := reader.Read(buffer)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("falha ao ler arquivo %s: %s", filename, err.Error())
		}

		fmt.Printf("[%s]", string(buffer[:position]))
	}
	return nil
}

func deletarArquivo(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("falha ao deletar arquivo %s: %s", filename, err.Error())
	}

	return nil
}
