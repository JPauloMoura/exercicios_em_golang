package main

import "fmt"

/*
	Funções anônimas são funções que não possuem nome e geralmente são usadas dentro de outra função.
*/

func main() {
	exemplo1()
	exemplo2()
	exemplo3()
}

// Aqui criamos a funcção e logo a executamos colocando "()" no final dela
func exemplo1() {
	func() { fmt.Println("exemplo 1") }()
}

// Podemos também passar paramentros para ela
func exemplo2() {
	mensagem := "exemplo 2"

	func(m string) { fmt.Println(m) }(mensagem)
}

// Outra coisa legal que podemos fazer é adicionar essas funcões dentro de variaveis
func exemplo3() {
	somar := func(n1, n2 int) int { return n1 + n2 }

	fmt.Println(somar(10, 2))
	fmt.Println(somar(111, 100))
	fmt.Println(somar(5, 5))
}
