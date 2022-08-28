package main

import (
	"fmt"
	"reflect"
)

/*
	Ordene uma lista de numeros inteiros de forma crescente:
	Intere sobre a lista de numeros e se o numero atual for menor que o proximo numero do lista
	(arr[i] > arr[i+1]) esses numeros devem trocar de lugar.
	Faça isso até nenhum numero trocar de lugar.

	Desafio: implemente também uma ordenação decrescente
*/
func main() {
	lista := []int{11, 7, 11, 11, 10, 0, 9, 8, 7, 6, 5, 5, 4, 4, 3, 2, 1, 0, 0}
	var sort BubbleSort

	fmt.Println(sort.Ascending_1(lista))
	fmt.Println(sort.Ascending_2(lista))
	fmt.Println(sort.Descending(lista))
}

//BubbleSort é um algoritmo de ordenação que compara dois elementos adjacentes e os troca até que estejam na ordem pretendida.
type BubbleSort struct{}

func (b BubbleSort) Ascending_1(numbers []int) []int {
	for {
		old := append(make([]int, 0, len(numbers)), numbers...)

		for current := 0; current < len(numbers)-1; current++ {
			if next := current + 1; numbers[current] > numbers[next] {
				numbers[next], numbers[current] = numbers[current], numbers[next]
				continue
			}
		}
		if reflect.DeepEqual(numbers, old) {
			break
		}
	}
	return numbers
}

func (b BubbleSort) Ascending_2(numbers []int) []int {
	endIndex := len(numbers) - 1

	for i := 0; i < endIndex; i++ {
		for index, current := range numbers {
			if index == endIndex {
				break
			}

			if next := index + 1; current > numbers[next] {
				numbers[index] = numbers[next]
				numbers[next] = current
			}
		}
	}
	return numbers
}

func (b BubbleSort) Descending(numbers []int) []int {
	endIndex := len(numbers) - 1

	for i := 0; i < endIndex; i++ {
		for index, current := range numbers {
			if index == endIndex {
				break
			}

			if next := index + 1; current < numbers[next] {
				numbers[index] = numbers[next]
				numbers[next] = current
			}
		}
	}
	return numbers
}
