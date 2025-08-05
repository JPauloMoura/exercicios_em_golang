package main

import "fmt"

func main() {
	a := 1
	b := 2.0
	Print(a, b)

	// Definindo um novo tipo baseado em float64.
	type MyFloat float64
	var c MyFloat = 3.5

	Print(a, c)

	// Definindo um novo tipo baseado em int.
	// Isso causaria um erro, pois MyInt não é int ou ~float64.
	// type MyInt int
	// var d MyInt = 10
	// Print(d, b)

	// usando a função PrintBoleto
	PrintBoleto(BoletoCaixa{Valor: 100})
	PrintBoleto(BoletoXP{Valor: 200, Pagador: "João"})

	// Testando comparação com generics
	fmt.Println(isEqual(1, 1))
	fmt.Println(isEqual(1, 1.1))
	fmt.Println(isEqual("teste", "teste 2"))
	fmt.Println(isEqual(true, false))
}

/*
Ao criar uma função genérica, você pode definir um conjunto de tipos aceitos usando o operador | (OU).
Isso permite que a função aceite múltiplos tipos, desde que eles estejam dentro do conjunto especificado.

No exemplo abaixo, a função Soma aceita tanto int quanto float64.
Note que usamos ~float64 para permitir tipos que são definidos como float64.
*/
func Print[T int | ~float64](a int, b T) {
	fmt.Println("Valor de a:", a)
	fmt.Println("Valor de b:", b)
}

/*
Utilizando constraint com tipos definidos pelo usuário.
Aqui, definimos uma constraint Number que aceita int e qualquer tipo definido como float64.
*/

type Number interface {
	~int | ~float64
}

// A função PrintNumber aceita qualquer tipo que satisfaça a constraint Number.
func PrintNumber[T Number](a T) {
	fmt.Println("Valor:", a)
}

/*
 Vamos usar constraints e generics para definir uma função que aceita diferentes tipos de boletos.
*/
type BoletoCaixa struct {
	Valor float64
}

type BoletoXP struct {
	Valor   float64
	Pagador string
}

type Boleto interface {
	BoletoCaixa | BoletoXP
}

func PrintBoleto[T Boleto](b T) {
	fmt.Println("Boleto:", b)

	// printa o tipo do boleto
	switch any(b).(type) {
	case BoletoCaixa:
		fmt.Println("Tipo do boleto: BoletoCaixa")
	case BoletoXP:
		fmt.Println("Tipo do boleto: BoletoXP")
	}
}

// Utilizando generics com a constraint comparable
// A constraint comparable permite que a função aceite qualquer tipo que possa ser comparado com ==.
// Isso inclui tipos primitivos como int, float64, string, e também tipos definidos pelo usuário que implementam a interface de comparação.
func isEqual[T comparable](a, b T) bool {
	return a == b
}
