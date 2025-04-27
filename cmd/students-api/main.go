package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MohakGupta2004/students-api/internal/config"
	"github.com/MohakGupta2004/students-api/internal/http/handlers/student"
)

func main() {
	cfg := config.MustHave()
	router := http.NewServeMux()

	router.HandleFunc("POST /api/student", student.New())

	server := &http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}

	// Handle OS signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in background
	go func() {
		log.Printf("Server is running at %s", cfg.HttpServer.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt
	<-done
	slog.Info("Shutting down server...")

	// Try graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
		os.Exit(1)
	}

	slog.Info("Server exited properly")
	os.Exit(0) // exit cleanly after shutdown
}
