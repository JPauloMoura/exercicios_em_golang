package main

import (
	"fmt"
)

type IFabricaEspotiva interface {
	criarSapato() ISapato
	criarCamisa() ICamisa
}

type ISapato interface {
	defineTamanho(pontuacao int)
	buscarTamanho() int
}

type ICamisa interface {
	defineCor(cor string)
	buscarCor() string
}

// buscarFabricaEspotiva é responsavel por devolver um tipo de fabricada especifica
// que cria itens da marca selecionada
func buscarFabricaEspotiva(marca string) (IFabricaEspotiva, error) {
	switch marca {
	case "adidas":
		return Adidas{}, nil
	case "nike":
		return Nike{}, nil
	}

	return nil, fmt.Errorf("nome de marca invalido: %s", marca)
}
