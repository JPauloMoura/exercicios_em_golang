package main

// Nike é responsavel por criar alguns itens como sapato e camisa.
// por implementar os metodos da inteface IFabricaEspotiva ela
// também se torna um "FabricaEspotiva"
type Nike struct{}

func (Nike) criarSapato() ISapato { return &NikeSapato{} }
func (Nike) criarCamisa() ICamisa { return &NikeCamisa{} }

type NikeSapato struct{ tamanho int }

func (s *NikeSapato) defineTamanho(tamanho int) { s.tamanho = tamanho }
func (s NikeSapato) buscarTamanho() int         { return s.tamanho }

type NikeCamisa struct{ cor string }

func (s *NikeCamisa) defineCor(cor string) { s.cor = cor }
func (s *NikeCamisa) buscarCor() string    { return s.cor }
