## Ninja nível 2
### 1. Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
```go
package main

import "fmt"

func main() {
	verDados(982749)
}

func verDados(n int) {
	fmt.Printf("Binário-> %b\nDecimal-> %d\nHexadecimal-> %#x", n, n, n)
}

```
#### Resultado:
```bash
Binário-> 11101111111011011101
Decimal-> 982749
Hexadecimal-> 0xefedd
```

[Resolução do Exercício](https://goplay.space/#GvxMOZjMBwx) 

<br/>

### 2. Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis. ( <, >, <=, >=, ==, != )
```go
package main

import "fmt"

func main() {
	maior := 10 > 6
	menor := 10 < 6
	maiorIgual := 10 >= 10
	menorIgual := 10 <= 10
	igual := "maria" == "maria"
	diferente := "maria" != "Maria"

	fmt.Printf(`
		10 > 6 ? %v
		10 < 6 ? %v
		10 >= 10 ? %v
		10 <= 10 ? %v
		"maria" == "maria" ? %v
		"maria" != "Maria" ? %v
	`, maior, menor, maiorIgual, menorIgual, igual, diferente)
}

```
#### Resultado:
```bash
10 > 6 ? true
10 < 6 ? false
10 >= 10 ? true
10 <= 10 ? true
"maria" == "maria" ? true
"maria" != "Maria" ? true
```
[Resolução do Exercício](https://goplay.space/#P7KNokpKFDD) 

<br/>

### 3. Crie constantes tipadas e não-tipadas e demonstre seus valores.
```go
package main

import "fmt"

const (
	semTipo     = 90.0
	comTipo int = 90
)

func main() {
	fmt.Printf("Sem tipagem: %v -> %T \n", semTipo, semTipo)
	fmt.Printf("Com tipagem: %d -> %T \n", comTipo, comTipo)
}
```
#### Resultado:
```bash
Sem tipagem: 90 -> float64 
Com tipagem: 90 -> int 
```
[Resolução do Exercício](https://goplay.space/#9Tjkqv3sDOR) 

<br/>

### 4. Crie um programa que:
- Atribua um valor int a uma variável
- Demonstre este valor em decimal, binário e hexadecimal
- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
- Demonstre esta outra variável em decimal, binário e hexadecimal
```go
package main

import "fmt"

func main() {
	v := 60
	fmt.Printf("| %d | %b | %#x |\n\n", v, v, v)

	v2 := v << 1
	fmt.Printf("| %d | %b | %#x |", v2, v2, v2)
}
```
#### Resultado:
```bash
| 60 | 111100 | 0x3c |

| 120 | 1111000 | 0x78 |
```
[Resolução do Exercício](https://goplay.space/#eAyX0VGuvwn) 

<br/>

### 5. Crie uma variável de tipo string utilizando uma raw string literal e demonstre-a.
```go
package main

import "fmt"

func main() {
	texto := `
		Lista de compras:
			[] Arroz
			[] Feijão
			[] Batata
	`
	fmt.Print(texto)
}
```
#### Resultado:
```bash
Lista de compras:
			[] Arroz
			[] Feijão
			[] Batata
```
[Resolução do Exercício](https://goplay.space/#veoAnlZETAu) 

<br/>

### 6. Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos demonstre estes valores.
```go
package main

import "fmt"

const (
	_ = iota + 2021
	ano2
	ano3
	ano4
	ano5
)

func main() {
	fmt.Print(ano2, ano3, ano4, ano5)
}
```
#### Resultado:
```bash
2022 2023 2024 2025
```
[Resolução do Exercício](https://goplay.space/#mCCv-MSLoqn) 

<br/>
