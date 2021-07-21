## Ninja nível 3
### 1. Exiba todos os números de 1-10000
```go
package main

import "fmt"

func main() {
	for n := 1; n <= 10000; n++ {
		fmt.Print(n, " ")
	}
}
```
#### Resultado:
```bash
1 2 3 4 5 6...
...9997 9998 9999 10000
```

[Resolução do Exercício](https://goplay.space/#h_dZGWmVh0O) 

<br/>

### 2. Exiba o unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
```go
package main

import "fmt"

func main() {
    for letra := 65; letra <= 90; letra++ {
        fmt.Printf("\n%d --> ", letra)
        for vezes := 1; vezes <= 3; vezes++ {
	    fmt.Printf("%#U ", letra)
	}
    }
}
```
#### Resultado:
```bash
65 --> U+0041 'A' U+0041 'A' U+0041 'A' 
66 --> U+0042 'B' U+0042 'B' U+0042 'B' 
...
89 --> U+0059 'Y' U+0059 'Y' U+0059 'Y' 
90 --> U+005A 'Z' U+005A 'Z' U+005A 'Z'
```

[Resolução do Exercício](https://goplay.space/#JVgE1fSVn5M) 

<br/>

### 3. Crie um loop utilizando a sintaxe: for condition {}. Utilize-o para demonstrar os anos desde que você nasceu
```go
package main

import "fmt"

func main() {
	cont := 1998

	for cont <= 2021 {
		fmt.Print(cont, " ")
		cont++
	}
}

```
#### Resultado:
```bash
1998 1999 2000 2001 2002 2003 2004 2005 2006 2007 2008 2009 2010 2011 2012 2013 2014 2015 2016 2017 2018 2019 2020 2021 
```

[Resolução do Exercício](https://goplay.space/#tNxA3VgUU5Q) 

<br/>

### 4. Crie um loop utilizando a sintaxe: for  {}. Utilize-o para demonstrar os anos desde que você nasceu
```go
package main

import "fmt"

func main() {
	nascimento := 1998
	atual := 2021

	for {
		if nascimento > atual {
			break
		}

		fmt.Print(nascimento, " ")
		nascimento++
	}
}
```
#### Resultado:
```bash
1998 1999 2000 2001 2002 2003 2004 2005 2006 2007 2008 2009 2010 2011 2012 2013 2014 2015 2016 2017 2018 2019 2020 2021 
```

[Resolução do Exercício](https://goplay.space/#kn2CxvRR92j) 

<br/>

### 5. Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
```go
package main

import "fmt"

func main() {
	for inicio := 10; inicio <= 100; inicio++ {
		fmt.Printf("%d/4 resta: %v\n", inicio, (inicio % 4))
	}
}

```
#### Resultado:
```bash
10/4 resta: 2
11/4 resta: 3
...
99/4 resta: 3
100/4 resta: 0
```

[Resolução do Exercício](https://goplay.space/#mdD5GVxby_3) 

<br/>

### 6. Crie um programa que demonstre o funcionamento da declaração if.
```go
package main

import "fmt"

func main() {
	fmt.Print("Bom dia! ")

	if eDomingo := false; eDomingo {
		fmt.Print("Hoje é Domingo!")
	}
}
```
#### Resultado:
```bash
Bom dia! 
```

[Resolução do Exercício](https://goplay.space/#h30GFb3qkRN) 

<br/>

### 7. Utilizando a solução anterior, adicione as opções else if e else.

```go
package main

import "fmt"

func main() {
	if eDomingo := true; eDomingo {
		fmt.Print("Bom dia! Hoje é Domingo dia de descansar!")
	} else {
		fmt.Print("Bom dia! Tenho que trabalhar duro hoje!")
	}
}
```
#### Resultado:
```bash
Bom dia! Hoje é Domingo dia de descansar!
```

[Resolução do Exercício](https://goplay.space/#rotxWkgxh3g) 

<br/>

### 8. Crie um programa que utilize a declaração switch, sem switch statement especificado.

```go
package main

import "fmt"

func main() {
	turno := "tarde"
	switch {
	case turno == "manhã":
		fmt.Print("Bom dia!")

	case turno == "tarde":
		fmt.Print("Boa tarde!")

	case turno == "noite":
		fmt.Print("Boa noite")
	}
}
```
#### Resultado:
```bash
Boa tarde!
```

[Resolução do Exercício](https://goplay.space/#jXczMcdDwS_x) 

<br/>

### 8. Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

```go
package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "volei":
		fmt.Print("amo jogar volei!")

	case "basquete":
		fmt.Print("amo jogar basquete!")

	case "futebol":
		fmt.Print("amo jogar futebol!")
	}
}
```
#### Resultado:
```bash
amo jogar futebol!
```

[Resolução do Exercício](https://goplay.space/#BgjRA1Uih0T) 
//teste
<br/>