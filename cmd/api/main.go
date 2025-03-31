package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kevalsabhani/pharmatail-backend/internal/healthcheck"
)

func main() {

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/healthcheck", healthcheck.HealthcheckHandler)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
