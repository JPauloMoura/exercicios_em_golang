### 1. Crie um array que suporte 5 valores to tipo int, atribua valores aos seus índices, utilize range e demonstre os valores do array, Utilizando format printing demonstre o tipo do array..

```go
package main

import "fmt"

func main() {
	arr := [4]int{10, 20, 30, 40}

	for i, v := range arr {
		fmt.Printf("Indíce %d: %d\n", i, v)
	}
	fmt.Printf("Tipo do array: %T", arr)
}
```
#### Resultado:
```bash
Indíce 0: 10
Indíce 1: 20
Indíce 2: 30
Indíce 3: 40
Tipo do array: [4]int%  
```

[Resolução do Exercício](https://goplay.space/#YTmJOJMfZN-) 

<br/>

### 2. Crie uma Slice do tipo int, atribua 10 valores a ela, utilize range para demonstrar todos estes valores, e utilize format printing para demonstrar seu tipo.
```go
package main

import "fmt"

func main() {
	slice := []int{0, 10, 20, 30, 40, 50, 70, 80, 90, 100}
	for i, v := range slice {
		fmt.Printf("Indíce  %d: %d\n", i, v)
	}
	fmt.Printf("Tipo do slice: %T", slice)
}
```
#### Resultado:
```bash
Indíce  0: 0
Indíce  1: 10
Indíce  2: 20
Indíce  3: 30
Indíce  4: 40
Indíce  5: 50
Indíce  6: 70
Indíce  7: 80
Indíce  8: 90
Indíce  9: 100
Tipo do slice: []int  
```

[Resolução do Exercício](https://goplay.space/#rOc2aV1kJMh) 

<br/>

### 3. Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
- Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
- Do quinto ao último item do slice (incluindo o último item!)
- Do segundo ao sétimo item do slice (incluindo o sétimo item!)
- Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
- Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
```go
package main

import "fmt"

func main() {
	slice := []int{0, 10, 20, 30, 40, 50, 70, 80, 90, 100}
	fmt.Printf("item 1-3: %v\n", slice[:3])
	fmt.Printf("item 5-10: %v\n", slice[5:])
	fmt.Printf("item 2-7: %v\n", slice[2:7])
	fmt.Printf("item 3-9: %v\n", slice[3:9])
	//usando len()
	fmt.Printf("item 3-9: %v\n", slice[3:len(slice)-1])
}
```
#### Resultado:
```bash
item 1-3: [0 10 20]
item 5-10: [50 70 80 90 100]
item 2-7: [20 30 40 50 70]
item 3-9: [30 40 50 70 80 90]
item 3-9: [30 40 50 70 80 90]  
```

[Resolução do Exercício](https://goplay.space/#NH4fw_1IpNs)

<br/>

### 4. Começando com a seguinte slice: `x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}`
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
- y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x.
```go
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Printf("slice: %v\n", x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)
	fmt.Printf("slice: %v\n", x)
}
```
#### Resultado:
```bash
slice: [42 43 44 45 46 47 48 49 50 51 52 53 54 55]
slice: [42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60] 
```

[Resolução do Exercício](https://goplay.space/#thmWw13j-G1)

<br/>

### 5. Começando com a seguinte slice: `x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}`
- Utilizando slicing e append, crie uma slice y que contenha os valores:              [42, 43, 44, 48, 49, 50, 51]
```go
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)

	fmt.Printf("slice: %v\n", x)
}
```
#### Resultado:
```bash
slice: [42 43 44 48 49 50 51]
```

[Resolução do Exercício](https://goplay.space/#umhk4V-0kba)

<br/>