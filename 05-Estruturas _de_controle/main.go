package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("== Escolha uma opção ==")

	fmt.Println("P -> pizza")
	fmt.Println("H -> hamburger")
	fmt.Println("L -> Lasanha")
	fmt.Println("R -> refrigerante")

	var opcao string
	fmt.Scan(&opcao)

	switch strings.ToUpper(opcao) {
	case "P":
		fmt.Println("OK! Pizza.")
	case "H":
		fmt.Println("OK! Hamburger.")
	case "L":
		fmt.Println("OK! Lasanha.")
	case "R":
		fmt.Println("OK! Refrigerante.")
	default:
		fmt.Println("Opção invalida")
		os.Exit(-1) //fecha o programa com erro
	}

	os.Exit(0) //fecha o programa normalmente

}
