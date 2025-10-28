package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/searchCep/internal"
	"net/http"
)

func main() {
	handler := internal.NewCepHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/cep/{cep}", handler.Execute)

	http.ListenAndServe(":8080", r)
}
