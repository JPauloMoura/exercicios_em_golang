package main

import (
	"fmt"
	"time"
)

/*
Vamos usar o padrão de worker pool para realizar o processo de envio de email
para uma lista contatos.
*/
func main() {
	emails := []string{
		"jp",
		"vick",
		"billie",
		"kiki",
	}
	sendEmails(emails, 3)
}

func sendEmails(emails []string, workers int) {
	toProcess := make(chan string, len(emails))
	done := make(chan bool, workers)

	for i := 0; i < workers; i++ {
		go worker(toProcess, done)
	}

	for _, email := range emails {
		toProcess <- email
	}
	close(toProcess)

	for i := 0; i < workers; i++ {
		<-done
	}

}

func worker(toProcess chan string, done chan bool) {
	for email := range toProcess {
		time.Sleep(time.Second)
		fmt.Println("enviado: ", email)
	}

	done <- true
}
