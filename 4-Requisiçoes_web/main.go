package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		fmt.Println("=== Informe a URL do site ===")
		var site string
		var sair string

		fmt.Scan(&site)

		msg, err := verificarSite(site)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(msg)

		fmt.Println("[0] Sair")
		fmt.Println("[1] Nova consulta")
		fmt.Scan(&sair)

		if sair == "0" {
			os.Exit(0)
		}
	}
}

func verificarSite(site string) (string, error) {
	resp, err := http.Get(site)
	var msg string

	if err != nil {
		msgError := fmt.Sprint("Não foi possivel verificar o site. Erro: ", err.Error())
		return "", errors.New(msgError)
	}

	if resp.StatusCode != http.StatusOK {
		msg = fmt.Sprint("Erro na requisição, statuscode: ", resp.StatusCode)
		return msg, nil
	}

	msg = fmt.Sprint("Site funcionando, statuscode: ", resp.StatusCode)
	return msg, nil
}
