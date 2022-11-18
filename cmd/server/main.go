package main

import (
	"net/http"

	"github.com/MatThHeuss/si_2020_2_api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/users", handlers.CreateUser)

	http.ListenAndServe(":8000", r)
}
