package domain

import (
	"context"

	"github.com/stephen776/go-template/internal/postgres"
)

// UserService defines CRUD for User model
type UserService interface {
	GetUsers(ctx context.Context) ([]postgres.User, error)
	// GetUserById(id int) (User, error)
	// CreateUser(u User) error
	// DeleteUser(id int) error
}