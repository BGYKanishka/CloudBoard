package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloudboard/internal/database"
	"cloudboard/internal/handlers"
	"cloudboard/internal/health"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env if it exists
	godotenv.Load("../../.env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize Database
	if err := database.InitDB(); err != nil {
		log.Printf("Warning: Database initialization failed: %v", err)
		log.Println("Server will still start to serve /health, but /ready and app routes will fail.")
	} else {
		defer database.CloseDB()
	}

	// Setup Router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health and Readiness
	r.Get("/health", health.HealthHandler)
	r.Get("/ready", health.ReadyHandler)

	// API Routes
	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", handlers.RegisterHandler)
			r.Post("/login", handlers.LoginHandler)
			r.Post("/logout", handlers.LogoutHandler)
			r.Get("/me", handlers.MeHandler)
		})

		r.Route("/tasks", func(r chi.Router) {
			// Require auth middleware later
			r.Get("/", handlers.ListTasksHandler)
			r.Post("/", handlers.CreateTaskHandler)
			r.Get("/{id}", handlers.GetTaskHandler)
			r.Put("/{id}", handlers.UpdateTaskHandler)
			r.Delete("/{id}", handlers.DeleteTaskHandler)
			r.Patch("/{id}/status", handlers.UpdateTaskStatusHandler)
		})

		r.Post("/profile/image", handlers.UploadProfileImageHandler)
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Graceful Shutdown
	go func() {
		log.Printf("Starting server on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
