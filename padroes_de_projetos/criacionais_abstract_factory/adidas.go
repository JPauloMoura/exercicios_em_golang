package main

// Adidas é responsavel por criar alguns itens como sapato e camisa.
// por implementar os metodos da inteface IFabricaEspotiva ela
// também se torna um "FabricaEspotiva"
type Adidas struct{}

func (Adidas) criarSapato() ISapato { return &AdidasSapato{} }
func (Adidas) criarCamisa() ICamisa { return &AdidasCamisa{} }

type AdidasSapato struct{ tamanho int }

func (s *AdidasSapato) defineTamanho(tamanho int) { s.tamanho = tamanho }
func (s AdidasSapato) buscarTamanho() int         { return s.tamanho }

type AdidasCamisa struct{ cor string }

func (s *AdidasCamisa) defineCor(cor string) { s.cor = cor }
func (s *AdidasCamisa) buscarCor() string    { return s.cor }
