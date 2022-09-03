package main

import "fmt"

/*
	Funções recursivas são funções que chamam elas mesmas para que tenhamos um resultado esperado
	quando uma condisão de parada é atendida.
*/
func main() {
	contadorRecursivo(10)
	fmt.Println("fatorial ->", factorial_calc(5))
}

// contadorRecursivo realiza uma contagem decresente chamando ele mesmo, mas a cada nova chamada ele diminui numero.
// faz isso enquanto o numero for maior que zero. Essa é sua condição de parada
func contadorRecursivo(numero int) {
	if numero > 0 {
		fmt.Println(numero)
		contadorRecursivo(numero - 1)
	}

	return
}

// Um exemplo classico do uso de recursividade é o calculo de fatoriais.
// onde  N! = n *(n-1)... 1
// ex1: 3! = 3 * 2 * 1 = 6
// ex2: 5! = 5 * 4 * 3 * 2 * 1 = 120
func factorial_calc(number int) int {
	if number == 0 || number == 1 {
		return 1
	}

	return number * factorial_calc(number-1)
}
