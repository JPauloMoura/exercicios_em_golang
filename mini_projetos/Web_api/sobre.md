1) Criamos um servido usando o pacode `http`:
```go
func main() {
	http.HandleFunc("/", handlePageWeb)
	http.ListenAndServe(":3333", nil)
}

func handlePageWeb(w http.ResponseWriter, r *http.Request) {
  
}

```

2) existe um site para buscarmos por libs do Go:
[godoc](https://godoc.org)

3) a lib do postgres so é usanda durante o tempo de execução do nosso programa
então para o vscode não deletar ela ao salvar o documento, usamos o underline na frente da importação do pkg.
`	_ "github.com/lib/pq"`