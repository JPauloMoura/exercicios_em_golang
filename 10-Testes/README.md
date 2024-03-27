## Execuções

Comando | Descrição
---     | ---
`$ go test` | Executa os testes que estão dentro do pacote atual
`$ go test ./...` | Executa todos os testes do projeto
`-v` | Utilize para ver mais detalhes dos testes "verbose"
`-cover` | Utilize para ver a porcetagem de cobertura dos testes
`$go test --coverprofile coverage.txt` | Gera um arquivo com as informações de cobertura
`$go tool cover --func=coverage.txt` | Interpreta as informações do arquivo e mostra no terminal
`$go tool cover --html=coverage.txt` | Mostra as linhas não cobertas em uma página web

