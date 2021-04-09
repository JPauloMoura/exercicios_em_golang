package main

import "fmt"

type Pessoa struct {
	Nome         string
	Idade        int
	TemGraduacao bool
	Endereco     Endereco
}

type Endereco struct {
	Cidade string
	Estado string
	Rua    string
	Numero int
}

func main() {
	// declaração 1
	var jose Pessoa = Pessoa{}
	jose.Nome = "José da Silva"
	jose.Idade = 18
	jose.Endereco.Rua = "Violeta"
	//observe que os valores não declarados receberam um valor padrão
	fmt.Println(jose) //{José da Silva 18 false {  Violeta 0}}

	// declaração 2
	carlos := Pessoa{}
	carlos.Nome = "Carlos da Silva"
	carlos.Idade = 18
	carlos.Endereco.Rua = "Violeta"
	fmt.Println(carlos)

	// declaração 3
	bruna := Pessoa{
		Nome: "Bruna da Silva",
		Endereco: Endereco{
			Numero: 786,
		},
	}
	fmt.Println(bruna)

	// declaração 4
	magali := Pessoa{"Magali da Silva", 18, true, Endereco{"Altos", "SP", "rua 2", 909}}
	fmt.Println(magali)

	//declaração 5 com ponteiros
	var fernanda *Pessoa
	fernanda = new(Pessoa)
	fmt.Println(fernanda)

}
