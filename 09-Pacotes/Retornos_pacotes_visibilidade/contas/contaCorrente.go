package contas

import "fmt"

// a visibilidade dos atributos, metodos e structs e definido pela
// primeira letra da palava.
// Maiusculo -> public, minusculo -> private

type ContaCorrente struct {
	Numero int
	Saldo  float64
}

func (c *ContaCorrente) Depositar(valor float64) {
	operacaoValida := valor > 0
	if operacaoValida {
		c.Saldo += valor
		fmt.Println("Novo Saldo: R$", c.Saldo)
		return
	}

	fmt.Println("Operação invalida")
}

func (c *ContaCorrente) Transferi(valor float64, destino *ContaCorrente) {
	operacaoValida := valor > 0 && c.Saldo >= valor
	if operacaoValida {
		c.Saldo -= valor
		destino.Saldo += valor
		fmt.Println("Sucesso")
		return
	}

	fmt.Println("Operação invalida")
}
