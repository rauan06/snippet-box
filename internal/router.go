package internal

import (
	"net/http"
	"snippet-box/internal/handler"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/snippet/view", handler.SnippetView)
	mux.HandleFunc("/snippet/create", handler.SnippetCreate)

	return mux
}
