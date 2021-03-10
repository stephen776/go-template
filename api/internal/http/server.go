package http

import (
	"context"
	"net/http"
	"time"

	"bitbucket.org/truefit/tf-manifest/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ShutdownTimeout is the time given for outstanding requests to finish before shutdown.
const ShutdownTimeout = 1 * time.Second

// Server represents an HTTP server and wraps all HTTP functionality
type Server struct {
	router chi.Router
	server *http.Server

	// services
	UserService models.UserService
}

// NewServer returns a new instance of the Server Struct
func NewServer() *Server {
	s := &Server{
		router: chi.NewRouter(),
		server: &http.Server{},
	}

	// setup the base router for our Server
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	// TODO:
	// - serve embedded assets ?
	// - metrics?
	// - not found handler
	// - auth

	// testing only
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	// register other routes

	// TODO: Port!
	s.server.Addr = ":8080"
	s.server.Handler = s.router

	return s
}

// Open begins listening for requests
func (s *Server) Open() (err error) {
	// TODO: - Initialize our secure cookie with our encryption keys?
	// TODO: port?

	go s.server.ListenAndServe()

	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}