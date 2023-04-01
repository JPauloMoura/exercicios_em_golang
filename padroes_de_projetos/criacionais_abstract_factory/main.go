package main

import (
	"fmt"
)

func main() {
	criarPecasDaAdidas()
	criarPecasDaNike()

	_, err := buscarFabricaEspotiva("lupo")
	if err != nil {
		fmt.Println(err)
		return
	}

}

func criarPecasDaNike() {
	nike, err := buscarFabricaEspotiva("nike")
	if err != nil {
		fmt.Println(err)
		return
	}

	camisa := nike.criarCamisa()
	camisa.defineCor("Verde")
	fmt.Println(camisa.buscarCor())

	sapato := nike.criarSapato()
	sapato.defineTamanho(38)
	fmt.Println(sapato.buscarTamanho())
}

func criarPecasDaAdidas() {
	adidas, err := buscarFabricaEspotiva("adidas")
	if err != nil {
		fmt.Println(err)
		return
	}

	camisa := adidas.criarCamisa()
	camisa.defineCor("Branca")
	fmt.Println(camisa.buscarCor())

	sapato := adidas.criarSapato()
	sapato.defineTamanho(41)
	fmt.Println(sapato.buscarTamanho())
}
