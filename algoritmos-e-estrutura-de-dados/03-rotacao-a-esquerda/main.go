package main

import "fmt"

/*
	o Algoritmo de rotação, mudar os item de um array de posição.

ex:

	1 rotação -> [1 7 8 9] -> [7 8 9 1]
	2 rotação -> [1 7 8 9] -> [8 9 1 7]
	4 rotação -> [1 7 8 9] -> [1 7 8 9] // nada muda pois após 4 rotações voltamos a posição inicial
*/
func main() {
	fmt.Println(leftRotation([]int{1, 2, 3}, 3))
}

func leftRotation(arr []int, rotations int) []int {
	if rotations > len(arr) {
		rotations = rotations % len(arr)
	}

	rigth := arr[:rotations]
	left := arr[rotations:]

	left = append(left, rigth...)

	return left
}
