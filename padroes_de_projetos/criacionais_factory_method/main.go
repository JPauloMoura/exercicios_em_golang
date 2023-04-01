package main

import "fmt"

type Creater interface {
	documentFactory(name string) (Document, error)
	validatorFactory(Document) error
}

type Factory struct{}

// documentFactory é responsavel por criar tipos de documentos diferentes de acordo com o input passado
func (Factory) documentFactory(name string) (Document, error) {
	switch name {
	case "rg":
		return RG{}, nil
	case "cpf":
		return CPF{}, nil
	}
	return nil, fmt.Errorf(fmt.Sprintf("invalid argument: '%s'", name))
}

func (Factory) validatorFactory(d Document) error {
	if len(d.name()) == 2 {
		return fmt.Errorf(fmt.Sprintf("invalid doc name: '%s'", d.name()))
	}
	return nil
}

// Document é uma interface que representa um documento
type Document interface {
	name() string
}

type RG struct{}

func (RG) name() string { return "RG" }

type CPF struct{}

func (CPF) name() string { return "CPF" }

func main() {
	var fabrica Factory

	doc, err := fabrica.documentFactory("RG")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(doc.name())

	err = fabrica.validatorFactory(doc)
	if err != nil {
		fmt.Println(err)
		return
	}
}
