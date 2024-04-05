package main

import "fmt"

func main() {
	fmt.Println(order([]int{8, 3, 5, 1, 6, 7}))
}

// Ordena do maio para o menor; O(n^2)
// isso porque interamos "n" vezes o tamanho da lista
// e maus "n" vezes para cada item dessa lista

func order(list []int) []int {
	var newList []int

	for len(list) > 0 {
		item, index := biggerItem(list)
		newList = append(newList, item)

		/*
			remove o index 3 do array
			[1 2 3 {4} 5] -> i3 = 4

			arr[:3] -> [1 2 3]
			arr[3+1:] -> [5]
		*/
		p1 := list[:index]
		p2 := list[index+1:]

		list = append(p1, p2...)
	}

	return newList
}

func biggerItem(list []int) (int, int) {
	var bigger int
	var index int

	for i, v := range list {
		if v > bigger {
			bigger = v
			index = i
		}

	}

	return bigger, index
}
