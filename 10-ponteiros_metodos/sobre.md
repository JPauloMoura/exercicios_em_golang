1) com a ajudade de ponteiros e endereços de memoria
faça a comparação entre duas variaveis de uma mesmo struct

2) usamos os ponteiros para nos referir ao conteudo de um endereço de memoria
-`*` conteudo de um endereço
-`&` endereço de memoria

3) Crie metodos para uma struct criada
dica: podemos passar multiplos argumetos para um metodo ou função usando `...`
ex:
```go
func (ms *myStruct) Metodo(nome string, ...int){

}
```

4) Quando recebemos o argumento de uma funçao estamos recebendo uma copia do conteudo enviado
com um outro endereço de memoria, e caso esse valor seja modificado ele não modificara o valor original
pois esse possui um outro endereço.
Ex:
```go
func main(){
  original := true
  resp := alterar(original)
  fmt.Println(resp) //true
}

func alterar(copia bool) bool{
  copia = false
  return copia
}
```
Para que funcione corretamente, devemos passar o endereço de memoria da nossa variavel,
assim o conteúdo dela sera alterado da forma esperada.
```go


```