package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/slumhunden/go_tutorials/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(r chi.Router) {
	router.Use(middleware.Authorization)
	
	router.Get("/coins", GetCoinBalance)
	})
}