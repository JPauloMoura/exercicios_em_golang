package main

import (
	"fmt"
	"reflect"
	"time"
)

const delay = 1

func main() {
	visualizarArray()
	fmt.Println("---------------------")
	visualizarSlice()
}

func visualizarArray() {
	var idades [5]int
	idades[0] = 10
	idades[1] = 13
	idades[2] = 15
	idades[3] = 16
	// idades[5] tera um valor padrão de int -> 0

	fmt.Println("Idades ->", idades)
	fmt.Println("Quantidade ->", len(idades))
	fmt.Println("capacidade do array ->", cap(idades))
	fmt.Println("Type ->", reflect.TypeOf(idades))

	for index, value := range idades {
		fmt.Println("index: ", index, "-> ", value)
		time.Sleep(delay * time.Second)
	}

}

func visualizarSlice() {
	idades := []int{10, 13, 15, 16, 19}

	fmt.Println("Idades ->", idades)
	fmt.Println("Quantidade ->", len(idades))
	fmt.Println("capacidade do slice ->", cap(idades))
	fmt.Println("Type ->", reflect.TypeOf(idades))

	for i, v := range idades {
		fmt.Println("index: ", i, "-> ", v)
		time.Sleep(delay * time.Second)
	}

	fmt.Println("Adicionado mais um item no slice")
	idades = append(idades, 20)

	fmt.Println("Idades ->", idades)
	fmt.Println("Quantidade ->", len(idades))
	fmt.Println("capacidade do slice ->", cap(idades))

}
