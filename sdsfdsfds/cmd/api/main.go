package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/your-org/your-app/internal/delivery/http/router"
	"github.com/your-org/your-app/internal/repository/memory"
	"github.com/your-org/your-app/internal/usecase"
)

func main() {
	// Wiring: Frameworks & Drivers (main) instantiates adapters and use case.
	repo := memory.NewRepository()
	uc := usecase.New(repo)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router.New(uc),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown: %v", err)
	}

	log.Println("server stopped")
}
