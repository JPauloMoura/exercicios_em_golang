package main

import (
	"fmt"
	"time"
)

/*
Gouroutines é uma forma que o Go nos dar de excutarmos ações de forma concorrente dentro de uma thead principal.
*/
func main() {
	// As gouroutines são finalizadas quando a thead principal(main) finaliza
	go fazerBolo()
	fazerCafe()

	// nesse casso "fazerCafé()" que está na thread principal, é mais rápido que "fazerBolo"
	// assim a main termina antes da thead que foi gerada
	fmt.Println("Fim da main!")
}

func fazerBolo() {
	fmt.Println("Fazendo bolo...")
	esperar(time.Second * 5)
	fmt.Println("== Bolo feito! ==")
}

func esperar(segundos time.Duration) { time.Sleep(segundos) }

func fazerCafe() {
	fmt.Println("Fazendo café...")
	esperar(time.Second * 4)
	fmt.Println("== Café feito! ==")
}
