package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Addr string
}

func Server(conf Config) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/healhz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	r.Get("/readz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	r.Get("/pods", podsHandler)
	r.Get("/nodes", nodesHandler)

	http.ListenAndServe(conf.Addr, r)
}
