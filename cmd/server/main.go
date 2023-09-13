package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/viniciusmgaspar/desafio-multithreading/internal/infra/webServer/handlers"
)

func main() {
	handler := handlers.NewCepHandler()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/consulta-cep/{cep}", handler.GetCep)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":8000", router)
}
