## Time line de erros que tive durante a execução

1) o servido não sobe:
Ao executar o `go run` o ele não trava o terminal segurando o processo do servidor
SOLUÇÃO:
 >  Alterei a porta do servido pois a outra estava em uso.

<br/>

2) 2021/05/02 22:51:22 pq: password authentication failed for user "joaopaulo@moura"
SOLUÇÃO:
> o string de conexão estava errada. errado: `use`, correto `user`

3) Ao entrar na página de edição do produto, o ID do produto era enviado como `""`
na request:
EX: `productID := r.URL.Query().Get("id") // productID --> ""`

SOLUÇÂO:
 >  não estava passando o id na query params
 ```html
    <a class="btn btn-info" href="/editar?{{.Id}}">Editar</a> <!--Errado -->
    <a class="btn btn-info" href="/editar?id={{.Id}}">Editar</a> <!--Correto -->
 ```