package main

import (
	"fmt"

	"github.com/JoaoPaulo/contas"
)

func main() {
	jose := contas.ContaCorrente{177, 300}
	bruna := contas.ContaCorrente{178, 100}
	fmt.Println(jose)
	fmt.Println(bruna)

	jose.Depositar(200)
	fmt.Println(jose)
	fmt.Println(bruna)

	jose.Transferi(500, &bruna) //é passado o endereço de bruna
	fmt.Println(jose)
	fmt.Println(bruna)
}
