package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/your-org/your-app/internal/entity"
	"github.com/your-org/your-app/internal/usecase"
)

// Handler groups HTTP controllers (Interface Adapters - controllers).
// Converts HTTP requests to use case calls and formats the response.
type Handler struct {
	uc *usecase.UseCase
}

// New creates the handler with the use case injected.
func New(uc *usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

// Health returns the API status.
func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

// GetEntity returns an entity by ID.
func (h *Handler) GetEntity(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "missing id")
		return
	}

	e, err := h.uc.GetEntity(r.Context(), id)
	if err != nil {
		if err == entity.ErrNotFound {
			writeError(w, http.StatusNotFound, "not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(e)
}

// CreateEntity creates an entity (body: {"id": "..."}).
func (h *Handler) CreateEntity(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body")
		return
	}
	if input.ID == "" {
		writeError(w, http.StatusBadRequest, "id required")
		return
	}

	e := &entity.Entity{ID: input.ID}
	if err := h.uc.CreateEntity(r.Context(), e); err != nil {
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(e)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
