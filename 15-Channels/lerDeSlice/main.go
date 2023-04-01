package main

import "fmt"

/*
Nesse exemplo temos uma lista de nomes que deve ser lidas em uma
gorountina e enviadas de volta para a thread main atravez
de um canal.
*/
func main() {
	pessoas := []string{"jp", "vick", "lucas", "maria", "fabio", "luiz", "igor"}

	ch := make(chan string)
	go ler(pessoas, ch)

	// dessa forma temos um controle de quando um canal deve ser fechado
	ouvirCanal(ch)
	//ou  simplesmente
	ouvirCanalSimplificado(ch)
}

func ler(pessoas []string, ch chan string) {
	for _, p := range pessoas {
		ch <- p
	}

	close(ch)
}

func ouvirCanal(ch chan string) {
	for {
		if msg, aberto := <-ch; aberto {
			fmt.Println(msg)
			continue
		}

		break
	}
}

// dessa forma continuamos recebendo mensagens de um canal enquanto ele
// ainda estiver aberto
func ouvirCanalSimplificado(ch chan string) {
	for msg := range ch {
		fmt.Println(msg)
	}
}
