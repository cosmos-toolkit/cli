package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/your-org/your-app/internal/delivery/http/handler"
	"github.com/your-org/your-app/internal/usecase"
)

// New returns an http.Handler with the API routes (chi router).
func New(uc *usecase.UseCase) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := handler.New(uc)

	r.Get("/health", h.Health)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/entities/{id}", h.GetEntity)
		r.Post("/entities", h.CreateEntity)
	})

	return r
}
