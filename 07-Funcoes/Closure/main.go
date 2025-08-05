package main

import "fmt"

/*
	Funções Closure, são funções que "lembram" do escopo de onde foram declaradas.

*/

func main() {
	listaDeNomes := []string{"jp", "vick", "billie", "rafael", "gustavo", "maria"}
	proximoDalista := PercorrerLista(listaDeNomes)

	fmt.Println(proximoDalista())
	fmt.Println(proximoDalista())
	fmt.Println(proximoDalista())
	fmt.Println(proximoDalista())
	fmt.Println(proximoDalista())
	fmt.Println(proximoDalista())
}

type ProximoItem func() string

// PercorrerLista devolve a função ProximoItem, que vai ter na sua "memória" o valor de todas as variáveis
// que foram declaradas no escopo em que ela foi criada.
// ou seja, a cada nova chamada da função ProximoItem que foi retornada, ela vai lembrar de todos os seu valores anteriores.
func PercorrerLista(lista []string) ProximoItem {
	var index int
	var item string

	funcBuscarItem := func() string {
		if index < len(lista) {
			item = lista[index]
			index++
		}
		return item
	}

	return funcBuscarItem
}
