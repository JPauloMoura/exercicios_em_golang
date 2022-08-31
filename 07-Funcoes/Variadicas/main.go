package main

import "fmt"

/*
	Funções Variadicas são funções que podem receber 0 ou N parametros de entrada de um mesmo tipo.
	Os paramentros devem ser sempre do mesmo e não podemos ter mais de um como entrada de uma função
*/
func main() {
	simple(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	multiplosArgumentos("olá", 1, 2, 3, 4, 5, 56, 6, 7)

	lista := []int{1, 2, 3, 4, 5}
	passandoLista(lista...)

	passandoQualquerTipo("1", 1, []bool{true, false})
}

func simple(numeros ...int) {
	fmt.Println(numeros) //[1 2 3 4 5 6 7 8 9 10]
}

func multiplosArgumentos(msg string, numeros ...int) {
	fmt.Println(numeros) //[1 2 3 4 5 56 6 7]
	fmt.Println(msg)     // olá
}

func passandoLista(numeros ...int) {
	fmt.Println(numeros) //[1 2 3 4 5]
}

func passandoQualquerTipo(any ...interface{}) {
	fmt.Println(any) //[1 1 [true false]]
}
