package main

import (
	"fmt"
	"time"
)

/*
Utilizamos o select para receber dados de vários canais, sem que um canal acabe boqueado o outro.
Imagine que temos uma lanchonete que uma xicará de café a cada 2 segundos
e os bolos são feitos a cada 4 segundos. A lanchonete faz esses item de forma concorrente
é não precisa espera um item para que o outro seja feito. Vamos ver como isso é feito usando canais.
*/
func main() {
	chCafe := make(chan string)
	chBolo := make(chan string)

	go fazerCafe(chCafe)
	go fazerBolo(chBolo)

	BuscarBoloECafeComSelect(chCafe, chBolo)
}

func fazerCafe(ch chan string) {
	for {
		time.Sleep(time.Second * 1)
		ch <- fmt.Sprintf("café, ")
	}
}

func fazerBolo(ch chan string) {
	for {
		time.Sleep(time.Second * 4)
		ch <- fmt.Sprintf("\n==> Bolo")
	}

}

// Se não usarmos o select quando estamos lendo varios canais, acabamos travando a leitura dos demais quando um deles é mais lento.
// No exemplo abaixo, temos que aguarda o bolo ser feito para que novamente termos outro café
func BuscarBoloECafe(chCafe, chBolo chan string) {
	var loop int
	for loop < 20 {
		fmt.Println(<-chCafe) // fazer café é rapido então logo liberamos esse canal
		fmt.Println(<-chBolo) // já fazer bolo demora mais, então vamo ficar travados aqui esperando o bolo

		loop++
	}
}

// Já utilizando o select, ficamos esperando dados dos dois canais ao mesmo tempo, printando a mensagem do
// que estiver pronto primeiro, sem esperá pelo outro.
// No exemplo sempre que o canal manda um café printamos eles, sem espera que o bolo esteje pronto
func BuscarBoloECafeComSelect(chCafe, chBolo chan string) {
	var loop int
	for loop < 20 {
		select {
		case cafe := <-chCafe:
			fmt.Print(cafe)

		case bolo := <-chBolo:
			fmt.Println(bolo)
		}

		loop++
	}
}
