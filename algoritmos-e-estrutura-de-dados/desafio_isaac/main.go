package main

import (
	"fmt"
	"strings"
)

/*
lista de feriados "01-01-2024, 25-01-2024, 29-03-2024, 21-04-2024, 01-05-2024"
parcelas "ID1:01-02-2024,ID2:05-03-2024,ID3:21-04-2024,ID4:01-05-2024,ID5:01-06-2024"

output: ID3, ID4
*/

func main() {
	vericarParcelasEmFereiado()
}

func vericarParcelasEmFereiado() {
	listaDeferiados := formatarListaDeferiados("01-01-2024,25-01-2024,29-03-2024,21-04-2024,01-05-2024")
	parcelas := formatarListaDeParcelas("ID1:01-02-2024,ID2:05-03-2024,ID3:21-04-2024,ID4:01-05-2024,ID5:01-06-2024")

	for _, feriado := range listaDeferiados {
		if id, sim := parcelas[feriado]; sim {
			fmt.Println(id)
		}
	}
}

// documentar funções
func formatarListaDeferiados(lista string) []string {
	return strings.Split(lista, ",")
}

func formatarListaDeParcelas(lista string) map[string]string {
	parcelas := strings.Split(lista, ",") // ID1:01-02-2024
	mapa := make(map[string]string)

	for _, parcela := range parcelas {
		parcelaComId := strings.Split(parcela, ":")

		id := parcelaComId[0]
		data := parcelaComId[1]
		mapa[data] = id
	}

	return mapa
}
