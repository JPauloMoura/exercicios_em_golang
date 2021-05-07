## Time line de erros que tive durante a execução

1) o servido não sobe:
Ao executar o `go run` o ele não trava o terminal segurando o processo do servidor
SOLUÇÃO:
 >  Alterei a porta do servido pois a outra estava em uso.

<br/>

2) 2021/05/02 22:51:22 pq: password authentication failed for user "joaopaulo@moura"
SOLUÇÃO:
> o string de conexão estava errada. errado: `use`, correto `user`
