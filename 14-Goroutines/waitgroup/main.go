package main

import (
	"fmt"
	"sync"
	"time"
)

/*
WaitGroup é uma das formas que temos para manipular e aguarda as gourotines.
*/
func main() {
	var wg sync.WaitGroup

	// aqui informamos que vamos ter que aguarda a finalização de 2 goroutines
	wg.Add(2)

	go func() {
		fazerBolo()
		wg.Done() // aqui informamos para o nosso grupo que a gouroutine finalizou
	}()

	go func() {
		fazerCafe()
		wg.Done() // aqui informamos para o nosso grupo que a gouroutine finalizou
	}()

	// aqui bloquiamos a thead principal até que todas as gouroutines do grupo finalize
	wg.Wait()
	fmt.Println("Fim da main!")
}

func fazerBolo() {
	for i := 1; i < 4; i++ {
		fmt.Println("*Bolo|passo", i)
		esperar()
	}
	fmt.Println("== Bolo feito! ==")
}

func esperar() { time.Sleep(time.Second * 1) }

func fazerCafe() {
	for i := 1; i < 4; i++ {
		fmt.Println("+Cafe|passo", i)
		esperar()
	}
	fmt.Println("== Café feito! ==")
}
