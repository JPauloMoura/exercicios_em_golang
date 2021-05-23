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