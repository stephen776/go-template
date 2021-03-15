package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"bitbucket.org/truefit/tf-manifest/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ShutdownTimeout is the time given for outstanding requests to finish before shutdown.
const ShutdownTimeout = 1 * time.Second

// Server represents an HTTP server and wraps all HTTP functionality
type Server struct {
	Addr string

	router chi.Router
	server *http.Server

	// services
	UserService domain.UserService
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

	// TODO: other middleware (jwt auth, etc...)

	// TODO:
	// - serve embedded assets ?
	// - not found handler
	// - auth
	// - request validation

	// register unauthenticated routes routes (auth routes for login, etc.)
	s.registerUserRoutes()

	// register authenticated routes

	s.server.Addr = s.Addr
	s.server.Handler = s.router

	return s
}

// Open begins listening for requests
func (s *Server) Open() (err error) {
	if s.Addr == "" {
		return fmt.Errorf("server port not set")
	}

	s.server.Addr = s.Addr

	go s.server.ListenAndServe()

	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}