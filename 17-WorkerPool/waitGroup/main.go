package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Vamos usar o padrão de worker pool para realizar o processo de envio de email
para uma lista contatos.

Nesse exemplo vamos usar apenas um channel e controlar as execuções usando wait group
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

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(toProcess, &wg)
	}

	for _, email := range emails {
		toProcess <- email
	}

	close(toProcess)

	wg.Wait()
}

func worker(toProcess chan string, wg *sync.WaitGroup) {
	for email := range toProcess {
		time.Sleep(time.Second)
		fmt.Println("email enviado para: ", email)
	}

	wg.Done()
}
