package main

import (
	"fmt"
	"time"
)

/*
O deadlock em channels acontece quando temos um canal esperando algum dado chegar
fazendo a thread aguarda esse dado, mas não temos ninguém pra enviar esse dado.
Ou seja, nosso programa vai ficar travado pra sempre.
*/
func main() {
	// channelControlado()
	channelComDeadLock()
}

// ============
func channelComDeadLock() {
	canal := make(chan string)
	go novaMSG(":( ", canal)

	for {
		// Quando nosso canal não recebe mais dados acontece o deadlock
		// fatal error: all goroutines are asleep - deadlock!
		fmt.Println("ch2 ", <-canal)
	}
}

func novaMSG(msg string, c chan string) {
	for i := 1; i < 4; i++ {
		time.Sleep(time.Second)
		c <- msg
	}
}

// =============

func channelControlado() {
	canal := make(chan string)
	go novaMSG2(":) ", canal)

	for {
		// Com o channel controlado, validamos o momento em que o channel está fechado
		// e paramos de ouvi-lo.
		if msg, aberto := <-canal; aberto {
			fmt.Println("ch1 ", msg)
			continue
		}

		break
	}
}

func novaMSG2(msg string, c chan string) {
	for i := 1; i < 4; i++ {
		time.Sleep(time.Second)
		c <- msg
	}

	// como não temos mais dados para enviar, fechamos o channel.
	close(c)
}
