package main

import "fmt"

/*
	A função init é uma função que é executada sempre no inicio do arquivo.
	Você pode criar uma função init por arquivo.
*/

var msg string

func main() {
	msg = "2-> main()"
	fmt.Println(msg)
}

func init() {
	msg = "1-> init()"
	fmt.Println(msg)
}
