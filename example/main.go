package main

import (
	"log"
	"net/http"

	"github.com/go-qbit/web-bootstrap4/example/handler"
)

func main() {
	log.Printf("Server listens on http://localhost:8000")
	if err := http.ListenAndServe("localhost:8000", handler.New()); err != nil {
		panic(err)
	}
}
