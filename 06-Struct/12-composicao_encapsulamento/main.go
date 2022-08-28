package main

import (
	"fmt"
)

type Carro struct {
	Modelo string
	Ano    int
}

func (c Carro) Andar() {
	fmt.Println("Carro", c.Modelo, "andando")
}

type Pessoa struct {
	Nome  string
	Carro Carro //composição
}

func main() {
	p := Pessoa{
		"Fabiana",
		Carro{"xpto", 2000},
	}

	fmt.Println(p.Nome)
	fmt.Println(p.Carro)
	p.Carro.Andar()

}
