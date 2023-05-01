package main

// IBuilder é a interface que definie os metodos de um builder
type IBuilderCasa interface {
	setTipoDeJanela()
	setTipoDePorta()
	setQtdDeAndares()
	getCasa() Casa
}

type builder string

var (
	Normal builder = "normal"
	Igloo  builder = "igloo"
)

// getBuilder retorna um builder apartir de um tipo
func getBuilder(tipoDeBuilder builder) IBuilderCasa {
	switch tipoDeBuilder {
	case Normal:
		return Casa{}.newBuilderNormal()

	case Igloo:
		return Casa{}.newBuilderIgloo()
	}

	return nil
}

// newDiretor cria uma nova estrutura apartir de um builder
func newDiretor(b IBuilderCasa) *Diretor {
	return &Diretor{b}
}

// Diretor é a estrutura responsavel por coordernar a criação de objetos
// apartir de um build
type Diretor struct {
	builder IBuilderCasa
}

func (d *Diretor) setBuilder(b IBuilderCasa) { d.builder = b }

// buildCasa realiza a construção do objeto casa
func (d *Diretor) buildCasa() Casa {
	d.builder.setTipoDePorta()
	d.builder.setTipoDeJanela()
	d.builder.setQtdDeAndares()

	return d.builder.getCasa()
}
