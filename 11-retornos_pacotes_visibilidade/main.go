package main

import (
	"fmt"

	"github.com/JoaoPaulo/contas"
)

func main() {
	jose := contas.ContaCorrente{Numero: 177, Saldo: 300}
	bruna := contas.ContaCorrente{Numero: 178, Saldo: 100}
	fmt.Println(jose)
	fmt.Println(bruna)

	jose.Depositar(200)
	fmt.Println(jose)
	fmt.Println(bruna)

	jose.Transferi(500, &bruna) //é passado o endereço de bruna
	fmt.Println(jose)
	fmt.Println(bruna)
}
