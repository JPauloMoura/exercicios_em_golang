1) Criamos um servido usando o pacode `http`:
```go
func main() {
	http.HandleFunc("/", handlePageWeb)
	http.ListenAndServe(":3333", nil)
}

func handlePageWeb(w http.ResponseWriter, r *http.Request) {
  
}

```