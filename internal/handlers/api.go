package handlers

import (
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    "github.com/ethandoes-code/go-api/internal/middleware"
)

// uppercase functions can pe used in other files,
// lowercase functions are private and only used in this file
func Handler(r *chi.Mux) {
    // Global middleware, Strip slashes ignores slashes
    r.Use(chimiddle.StripSlashes)

    // Set up route
    r.Route("/account", func(router chi.Router) {
        // Middleware for /account route
        router.Use(middleware.Authorization)

        router.Get("/coins", GetCoinBalance)
    })
}
