package websocket

import "net/http"

// PlayerServer  é a estrutura que representa o servidor de jogadores
type PlayerServer struct {
	Store PlayerStore
	http.Handler
}

type PlayerStore interface {
}
