package main

import (
	"fmt"
	"strings"
)

/*
Dado um numero inteiro de 1 a 3999, transforme esse numero em um algorimos romano
ex: 1 -> I,  119 -> CXIX
I	1
V	5
X	10
L	50
C	100
D	500
M	1000
*/
func main() {
	fmt.Println(toRoman(14))
}

func toRoman(number int) string {
	type Parse struct {
		value int
		slug  string
	}

	list := []Parse{
		{value: 1000, slug: "M"},
		{value: 900, slug: "CM"},
		{value: 500, slug: "D"},
		{value: 400, slug: "CD"},
		{value: 100, slug: "C"},
		{value: 90, slug: "XC"},
		{value: 50, slug: "L"},
		{value: 40, slug: "XL"},
		{value: 10, slug: "X"},
		{value: 9, slug: "IX"},
		{value: 5, slug: "V"},
		{value: 4, slug: "IV"},
		{value: 1, slug: "I"},
	}

	var roman string
	for _, p := range list {
		qtd := number / (p.value)

		if qtd <= 0 { // opcional pois as atribuições abaixo não teram efeito quando qtd for "0"
			continue
		}

		number -= qtd * p.value
		roman += strings.Repeat(p.slug, qtd)
	}

	return roman
}
