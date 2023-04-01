package main

import (
	"fmt"
)

func main() {
	canalNaMesmaThread()
	canalNaMesmaThreadComCapacidade()
}

func canalNaMesmaThread() {
	ch := make(chan string)
	// nesse momento que a msg é enviada, o programa para
	// já que não tem niguém ouvindo esse canal
	ch <- "oi"

	fmt.Println(<-ch)
}

func canalNaMesmaThreadComCapacidade() {
	// agora informamos que esse canal possui espaço de armazenamento
	// ou seja, ele vai conseguir guarda apenas 1 msg enquanto espera
	// alguém pra ler dele
	ch1 := make(chan string, 1)
	ch1 <- "ch1"
	fmt.Println(<-ch1)

	ch2 := make(chan string, 2)
	ch2 <- "ch2"
	ch2 <- "ch2"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// se passarmos uma msg a mais do que sua capacidade ela trava
	// por que no momento que passamos uma segunda mensagem, precisamos remover a outra
	// mas nesse momento não temos ninguém ouvindo.
	ch3 := make(chan string, 1)
	ch3 <- "ch3"
	ch3 <- "ch3" // trava
	fmt.Println(<-ch3)

	// a mesma coisa acontece se tentarmos ouvir mais de uma msg
	// msg essa que nunca foi envida
	ch4 := make(chan string, 1)
	ch4 <- "ch4"
	fmt.Println(<-ch4)
	fmt.Println(<-ch4) //trava
}
