package domain

import (
	"context"

	"bitbucket.org/truefit/tf-manifest/internal/store"
)

// UserService defines CRUD for User model
type UserService interface {
	GetUsers(ctx context.Context) ([]store.User, error)
	// GetUserById(id int) (User, error)
	// CreateUser(u User) error
	// DeleteUser(id int) error
}