package services

import (
	"context"

	"bitbucket.org/truefit/tf-manifest/internal/store"
)

// UserService is the implemation of the UserService domain interface
type UserService struct {
	q *store.Queries
}

// GetUsers returns all users
func (s *UserService) GetUsers(ctx context.Context) ([]store.User, error) {
	return s.q.GetUsers(ctx)
}

// NewUserService creates a new instance of the UserRepository
func NewUserService(q *store.Queries) *UserService {
	return &UserService{q}
}