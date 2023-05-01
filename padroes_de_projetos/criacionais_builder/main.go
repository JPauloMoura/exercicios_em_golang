package main

import (
	"fmt"
	"log"
)

func main() {
	bNormal := getBuilder("normal")
	if bNormal == nil {
		log.Fatal("builder normal e nil")
	}
	diretor := newDiretor(bNormal)
	casa := diretor.buildCasa()

	bIgloo := getBuilder("igloo")
	if bIgloo == nil {
		log.Fatal("builder igloo e nil")
	}
	diretor.setBuilder(bIgloo)
	igloo := diretor.buildCasa()

	fmt.Println(casa)
	fmt.Println(igloo)
}
