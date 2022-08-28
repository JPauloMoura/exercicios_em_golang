package main

import "fmt"

type Carro struct {
	Cor       string
	qtdPortas int
	Flex      bool
}

func (c *Carro) alterarCor(cor string) {
	c.Cor = cor
	fmt.Println(c.Cor)
}

func compararStructs() {
	//Compara os valores
	uno := Carro{"Azul", 4, false}
	ford := Carro{"Azul", 4, false}
	fmt.Println(uno == ford) //true

	//Compara os valores
	gol := Carro{"Preto", 2, false}
	corola := Carro{"Branco", 4, false}
	fmt.Println(gol == corola) //false

	//Compara o  endereço de memoria
	var combi *Carro
	combi = new(Carro)

	var toyota *Carro
	toyota = new(Carro)

	fmt.Println(combi == toyota) //false

	fmt.Println("combi:", combi, "Toyota:", toyota)   //conteudo
	fmt.Println("combi:", &combi, "Toyota:", &toyota) //endereço de memoria
}

func main() {
	// compararStructs()
	fmt.Println("================")

	uno := Carro{"Preto", 2, true}
	fmt.Println(uno.Cor)
	uno.alterarCor("Rosa")
}
