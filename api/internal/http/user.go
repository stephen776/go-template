package http

import (
	"encoding/json"
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
	users, err := s.UserService.GetUsers()

	if err != nil {
		w.Write([]byte("oops..."))
		return
	}

	// TODO: need a real response struct and real json encoding
	w.Header().Set("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			w.Write([]byte("oops...encoding error"))
			return
		}
}

