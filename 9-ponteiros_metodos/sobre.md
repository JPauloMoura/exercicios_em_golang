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