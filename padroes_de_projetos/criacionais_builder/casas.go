package main

// Casa é a strutura que define os atributos de uma casa
type Casa struct {
	tipoDeJanela string
	tipoDePorta  string
	andares      int
}

// newBuilderNormal retorna uma estrutura de builder para uma casa normal
func (c Casa) newBuilderNormal() *BuilderNormal {
	return &BuilderNormal{}
}

// newBuilderIgloo retorna uma estrutura de builder para uma casa normal
func (c Casa) newBuilderIgloo() *BuilderIgloo {
	return &BuilderIgloo{}
}

// BuilderNormal é a estrutura reponsavel por criar casas do tipo normal
type BuilderNormal struct {
	Casa
}

func (b *BuilderNormal) setTipoDeJanela() { b.tipoDeJanela = "Janela de vidro" }
func (b *BuilderNormal) setTipoDePorta()  { b.tipoDePorta = "Porta de madeira" }
func (b *BuilderNormal) setQtdDeAndares() { b.andares = 1 }
func (b BuilderNormal) getCasa() Casa     { return Casa{b.tipoDeJanela, b.tipoDePorta, b.andares} }

// BuilderIgloo é a estrutura reponsavel por criar casas do tipo igloo
type BuilderIgloo struct {
	Casa
}

func (b *BuilderIgloo) setTipoDeJanela() { b.tipoDeJanela = "Janela de gelo" }
func (b *BuilderIgloo) setTipoDePorta()  { b.tipoDePorta = "Porta de gelo" }
func (b *BuilderIgloo) setQtdDeAndares() { b.andares = 1 }
func (b BuilderIgloo) getCasa() Casa     { return Casa{b.tipoDeJanela, b.tipoDePorta, b.andares} }
