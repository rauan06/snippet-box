package main

import (
	"log"
	"net/http"
	"snippet-box/internal"
)

func main() {
	log.Println("starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", internal.Router()))
}
