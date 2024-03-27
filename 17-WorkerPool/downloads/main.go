package main

import (
	"fmt"
	"time"
)

/*
Imagine que você tem uma aplicação que precisa fazer o download de várias imagens de diferentes URLs.
Usar um Worker Pool pode melhorar o desempenho, permitindo que vários downloads ocorram simultaneamente.
*/

func main() {
	l := []string{
		"url-1",
		"url-2",
		"url-3",
		"url-4",
		"url-5",
		"url-6",
	}
	execDownloads(l, 4)
}

func execDownloads(urls []string, numProcess int) {
	// criamos o canal de jobs com buffer para que possamos
	// enviar todas as urls que precisamos, sem espera que elas sejam processadas.
	jobs := make(chan string, len(urls))
	completed := make(chan bool)

	// inicializar os workers em varias goroutines
	for i := 0; i < numProcess; i++ {
		fmt.Println("init worker ", i+1)
		go workerDownload(jobs, completed)
	}

	// envia as urls para dentro do canal,para que os workes comecem a pega-lás
	for _, url := range urls {
		jobs <- url
		fmt.Println("send url: ", url)
	}

	close(jobs)

	for i := 0; i < numProcess; i++ {
		<-completed
		fmt.Println("completed")
	}

	close(completed)
}

func workerDownload(jobs chan string, completed chan bool) {
	// ficamos recupenando os dados do canal, e quando acaba saimos do loop
	for url := range jobs {
		time.Sleep(time.Second * 3)
		fmt.Println("Download OK!", url)
	}

	completed <- true
}
