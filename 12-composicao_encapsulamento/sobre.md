1) No go não temos herança, mas usamos composição que é usar uma struct como tipo de outra
EX:

```go
type Pessoa struct{
  Nome string
}

type Diretor struct{
  Pessoa Pessoa
}
```
dessa forma cria uma struct conta correte em um pacote conta, e uma struct titular no pacote cliente.