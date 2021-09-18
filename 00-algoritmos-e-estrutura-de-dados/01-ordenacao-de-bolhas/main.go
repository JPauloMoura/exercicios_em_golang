package main

import "fmt"

/*
	Ordene uma lista de numeros inteiros de forma crescente:
	Intere sobre a lista de numeros e se o numero atual for menor que o proximo numero do lista
	(arr[i] > arr[i+1]) esses numeros devem trocar de lugar.
	Faça isso ate nenhum numero trocar de lugar.

	Desafio: implemente também uma ordenação decrescente
*/
func main() {
	numbers := sortAsc([]int{10, 2, 9, 11, 34})
	fmt.Println(numbers)
}

func sortAsc(numbers []int) []int {
	var end bool

	for {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] < numbers[i+1] {
				foo := numbers[i+1]
				numbers[i+1] = numbers[i]
				numbers[i] = foo
				continue
			}

			end = true
		}
		if end {
			break
		}
	}

	return numbers
}
