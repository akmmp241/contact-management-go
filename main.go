package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func main() {
	server := InitializedServer()

	log.Printf("Listening on %s", server.Addr)
	err := server.ListenAndServe()
	log.Fatal(err)
}
