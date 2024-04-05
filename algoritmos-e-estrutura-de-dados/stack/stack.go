package main

import "fmt"

type Pilha struct {
	valores []any
}

func (f *Pilha) Add(item any) {
	f.valores = append(f.valores, item)
}

func (f *Pilha) Remove() any {
	if len(f.valores) == 0 {
		return nil
	}

	lastItem := len(f.valores) - 1
	removed := f.valores[lastItem]
	f.valores = f.valores[:lastItem]
	return removed
}

func NovaPilha() *Pilha {
	return &Pilha{}
}

func main() {
	fila := NovaPilha()

	fila.Add(1)
	fila.Add("oi")
	fila.Add(0)
	fila.Add("fala")

	for i := 0; i < 4; i++ {
		fmt.Println(fila.valores...)
		fila.Remove()
	}

}
