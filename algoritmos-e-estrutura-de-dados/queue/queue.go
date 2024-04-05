package main

import "fmt"

type Fila struct {
	valores []any
}

func (f *Fila) Push(item any) {
	f.valores = append(f.valores, item)
}

func (f *Fila) Pop() any {
	if len(f.valores) == 0 {
		return nil
	}

	removed := f.valores[0]
	f.valores = f.valores[1:]
	return removed
}

func NovaFila() *Fila {
	return &Fila{}
}

func main() {
	fila := NovaFila()

	fila.Push(1)
	fila.Push("oi")
	fila.Push(0)
	fila.Push("fala")

	for i := 0; i < 4; i++ {
		fmt.Println(fila.valores...)
		fila.Pop()
	}

}
