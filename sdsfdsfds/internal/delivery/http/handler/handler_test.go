package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/your-org/your-app/internal/repository/memory"
	"github.com/your-org/your-app/internal/usecase"
)

func TestHealth(t *testing.T) {
	repo := memory.NewRepository()
	uc := usecase.New(repo)
	h := New(uc)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	h.Health(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Health: got status %d", rec.Code)
	}
}

func TestCreateEntity(t *testing.T) {
	repo := memory.NewRepository()
	uc := usecase.New(repo)
	h := New(uc)
	body := bytes.NewReader([]byte(`{"id":"e1"}`))
	req := httptest.NewRequest(http.MethodPost, "/api/v1/entities", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.CreateEntity(rec, req)
	if rec.Code != http.StatusCreated {
		t.Errorf("CreateEntity: got status %d, body %s", rec.Code, rec.Body.String())
	}
}
