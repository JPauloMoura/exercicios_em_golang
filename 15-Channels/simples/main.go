package main

import (
	"fmt"
	"time"
)

/*
Channels são canais utilizados pelo Go para receber e enviar dados.
Os dados enviados por eles devem ser do mesmos tipo, e por isso quando criamos um channel
indicamos qual o tipo de dado que irá se trafegado por eles.
Eles geralamente são usados para sincronizar goroutines.
*/
func main() {
	// criamos o channel e passamos ele pra dentro de uma goroutine
	canal := make(chan string)
	var count int

	// o loop infinito é usando aqui para ficar recebendo mensagens continuamente até
	// ter recebido 3 msgs
	for count < 3 {
		go NovaMsg("oi", canal)
		// Apartir de agora temos uma ponte entre a gorountine e a thread main

		// A main fica bloqueada aguardando ter algum conteúdo no canal
		fmt.Println("Aguardando o canal ter alguma msg...")

		msg := <-canal
		// Após receber o dado, a main da continuidade
		fmt.Println("Recebido:", msg)
		count++
	}
}

func NovaMsg(msg string, c chan string) {
	// Aqui criamos um delay para que a msg
	// demore algun segundo para entrar no canal
	for i := 1; i < 4; i++ {
		time.Sleep(time.Second)
		fmt.Println(i, ".. ")
	}

	// Adiciona uma msg ao canal
	c <- msg
}
