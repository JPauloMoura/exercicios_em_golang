package main

import "fmt"

// 1) uma interface é como um 'contrato' e define os metodos que uma struct deve possuir para ser consideranda desse tipo
// ex: se temos a interface A com os metodos b1 e b2, todo struct que implementar os metodos b1 e b2 será do tipo A

// 2) Uma interface é a definição de um conjunto de métodos comuns a um ou mais tipos. É o que permite a criação de tipos polimórficos em Go.

// 3) Com a ajuda de interface podemos criar funçoes que recebem como argumento um tipo de interface, e todos as struct
// que são do tipo dessa interface podem ser usadas.

// 4) funções que recebem uma interface podem ter comportamentos diferentes dependendo do seu tipo original
func main() {
	f := Funcionario{Nome: "Carlos", ativo: false}
	adm := Adm{Funcionario: Funcionario{Nome: "jp", ativo: true}, Permisao: "tudo"}
	f.Info()
	trocarStatusAtivo(&adm)
}

func trocarStatusAtivo(f IFuncionario) {
	f.setAtivo()
}

type IFuncionario interface {
	Info()
	setAtivo()
}

type Funcionario struct {
	Nome  string
	ativo bool
}

type Adm struct { // composição
	Funcionario
	Permisao string
}

type Cliente struct {
	Nome  string
	Email string
	ativo bool
}

func (f Funcionario) Info() {
	fmt.Println("Nome:", f.Nome)
	fmt.Println("Ativo:", f.ativo)
}

func (f *Adm) setAtivo() {
	f.ativo = !f.ativo
	f.Info()
}

func (c Cliente) Info() {
	fmt.Println("Cliente:", c.Nome)
	fmt.Println("Email:", c.Email)
	fmt.Println("Ativo:", c.ativo)
}
