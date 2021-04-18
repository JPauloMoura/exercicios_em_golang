package main

import "fmt"

func main() {
	nome := "Maria"
	fmt.Println("variavel original:", &nome, "-->", nome) //	0xc000010240--> [ Maria ]

	alterarNome1(nome)
	fmt.Println("nenhuma modificação:", &nome, "-->", nome) // 0xc000010260--> [ Maria ]

	alterarNome2(&nome)
	fmt.Println("alteração realizada:", &nome, "-->", nome) // 0xc000010260--> [ Maria ]
}

//aqui é criado uma copia do parâmetro enviado com o mesmo conteúdo
//mas com um outro entereço de memória
func alterarNome1(nome string) {
	fmt.Println("alterarNome1:", &nome, "-->", nome) //endereco: 0xc000010240--> [ Maria ]
	nome = "Roberta"
}

//aqui o parâmetro recebido é o endereço de memória da variável 'nome'
//ou seja um ponteiro. Assim modificamos o conteúdo daquele endereço
func alterarNome2(nome *string) {
	fmt.Println("alterarNome2:", nome, "-->", *nome) //endereco: 0xc000010240--> [ Maria ]
	*nome = "Roberta"
}
