package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// registerUserRoutes is a helper to setup routing for /users
func (s *Server) registerUserRoutes() {
	s.router.Route("/users", func(r chi.Router){

		// just to test...
		r.Get("/hello", s.handleSayHello)

		// list all users at /users
		r.Get("/", s.handleUsersIndex)
	})
}

func (s *Server) handleSayHello (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello /users!"))
}

func (s *Server) handleUsersIndex (w http.ResponseWriter, r *http.Request) {
	users, err := s.UserService.GetUsers(r.Context())

	if err != nil {
		w.Write([]byte("oops..."))
		return
	}

	if err := RenderJSON(w, 200, users); err != nil {
		w.Write([]byte("oops...encoding error"))
		return
	}
}

