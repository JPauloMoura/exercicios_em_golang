package main

import "fmt"

type FuncionarioIT interface {
	Info()
	setAtivo()
}

type Funcionario struct {
	Nome  string
	Email string
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
	fmt.Println("Email:", f.Email)
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

func trocarStatusAtivo(f FuncionarioIT) {
	f.setAtivo()
}

func main() {
	f := Funcionario{Nome: "Carlos", Email: "c@gmail", ativo: false}
	adm := Adm{Funcionario: Funcionario{Nome: "jp", Email: "j@gmail", ativo: true}, Permisao: "tudo"}
	f.Info()
	trocarStatusAtivo(&adm)

}
