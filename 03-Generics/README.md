# Go Generics: Uma Visão Geral

Go 1.18 introduziu uma das funcionalidades mais esperadas pela comunidade: os **parâmetros de tipo**, mais conhecidos como **Generics**. Essa adição permite a criação de funções e tipos que podem operar em um conjunto de tipos diferentes, aumentando a reutilização de código e a segurança de tipos.

Este documento explora os conceitos fundamentais de Generics em Go, com base nos exemplos práticos encontrados em `main.go`.

## 1. Funções Genéricas Básicas

É possível criar funções que aceitam múltiplos tipos usando o operador `|` (OU) na declaração do parâmetro de tipo.

O caractere `~` (til) é usado para indicar que queremos aceitar não apenas o tipo exato, mas também qualquer tipo que tenha ele como seu tipo subjacente.

**Exemplo:**

A função `Print` abaixo aceita um `int` e um segundo parâmetro `T` que pode ser `int` ou qualquer tipo que tenha `float64` como seu tipo subjacente (como `type MyFloat float64`).

```go
// Definindo um novo tipo baseado em float64.
type MyFloat float64

func Print[T int | ~float64](a int, b T) {
	fmt.Println("Valor de a:", a)
	fmt.Println("Valor de b:", b)
}

func main() {
    // Uso com float64
	Print(1, 2.0)

    // Uso com MyFloat
    var c MyFloat = 3.5
	Print(1, c)
}
```

## 2. Constraints (Restrições)

Para evitar a repetição de listas de tipos e criar restrições mais claras e reutilizáveis, podemos usar interfaces como *constraints*. Uma *constraint* define o conjunto de tipos que um parâmetro de tipo pode aceitar.

**Exemplo:**

Aqui, criamos a *constraint* `Number`, que permite tipos com base em `int` ou `float64`. A função `PrintNumber` usa essa *constraint* para ser mais concisa e legível.

```go
// A interface Number atua como uma constraint.
type Number interface {
	~int | ~float64
}

// A função aceita qualquer tipo que satisfaça a constraint Number.
func PrintNumber[T Number](a T) {
	fmt.Println("Valor:", a)
}
```

## 3. Generics com Tipos Estruturados (Structs)

Generics são especialmente úteis para criar funções que operam sobre diferentes tipos de `structs` que compartilham uma lógica comum, sem a necessidade de usar interfaces vazias (`interface{}`) e asserções de tipo.

**Exemplo:**

Temos dois tipos de boletos, `BoletoCaixa` e `BoletoXP`. Criamos uma *constraint* `Boleto` que une os dois tipos. A função `PrintBoleto` pode, então, aceitar qualquer um dos tipos de boleto de forma segura.

```go
type BoletoCaixa struct {
	Valor float64
}

type BoletoXP struct {
	Valor   float64
	Pagador string
}

// A constraint Boleto permite BoletoCaixa ou BoletoXP.
type Boleto interface {
	BoletoCaixa | BoletoXP
}

// Função genérica que imprime qualquer tipo de boleto.
func PrintBoleto[T Boleto](b T) {
	fmt.Println("Boleto:", b)

	// Opcionalmente, podemos verificar o tipo específico em tempo de execução.
	switch any(b).(type) {
	case BoletoCaixa:
		fmt.Println("Tipo do boleto: BoletoCaixa")
	case BoletoXP:
		fmt.Println("Tipo do boleto: BoletoXP")
	}
}

func main() {
    PrintBoleto(BoletoCaixa{Valor: 100})
	PrintBoleto(BoletoXP{Valor: 200, Pagador: "João"})
}
```

## 4. A Constraint `comparable`

Go fornece uma *constraint* pré-definida chamada `comparable`. Ela pode ser usada com qualquer tipo que suporte os operadores de comparação `==` e `!=`, como `int`, `string`, `bool`, ponteiros, canais e structs compostas por tipos comparáveis.

**Exemplo:**

A função `isEqual` utiliza a *constraint* `comparable` para comparar dois valores do mesmo tipo.

```go
// A função aceita qualquer tipo que possa ser comparado.
func isEqual[T comparable](a, b T) bool {
	return a == b
}

func main() {
    fmt.Println(isEqual(1, 1))           // true
    fmt.Println(isEqual("a", "b"))       // false
    fmt.Println(isEqual(true, true))     // true
}
```


