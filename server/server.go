package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/rbonnat/blockchain-in-go/server/controller"
	"github.com/rbonnat/blockchain-in-go/service"
)

// Run Initialize router and launch http server
func Run(ctx context.Context, port string) error {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initialize services
	s := service.New(ctx, time.Now)

	// Initialize routes
	r.Get("/", controller.HandleGetBlockchain(s))
	r.Post("/", controller.HandleWriteBlock(s))

	// Launch http server
	log.Printf("Server listening on port: '%s'", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Printf("Cannot launch http server: '%v'", err)
	}

	return err
}
