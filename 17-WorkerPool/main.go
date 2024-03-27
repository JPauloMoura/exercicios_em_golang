package main

import (
	"fmt"
	"time"
)

/*
Worker pool é um padrão que usamos para executar uma tarefa mais rapidamente.
Utilizamos channels e goroutines para processar a tarefa dividindo o trabalho em varios workers
*/
func main() {
	start := time.Now()
	fazerCafeWoker(3)

	t := time.Since(start)
	fmt.Println("duração:", t)
}

func fazerCafeSimples(qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Println("feito!")
	}
}

func fazerCafeWoker(qtd int) {
	input := make(chan int)
	output := make(chan int)

	for i := 0; i < qtd; i++ {
		go fazCafe(input, output)
	}

	for i := 0; i <= qtd; i++ {
		input <- i + 1
	}
	close(input)

	for v := range output {
		fmt.Println(v)
	}
}

func fazCafe(in, out chan int) {
	for v := range in {
		time.Sleep(time.Second)
		out <- v
		fmt.Println("feito: ", v)
	}
}
