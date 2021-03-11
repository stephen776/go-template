package store

import (
	"bitbucket.org/truefit/tf-manifest/pkg/models"
)

// UserService performs DB operations against the users table
type UserService struct {
	db *DB
}

// GetUsers returns all users
func (s *UserService) GetUsers() ([]*models.User, error) {
	users := make([]*models.User, 0)

	err := s.db.client.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// NewUserService creates a new instance of the UserRepository
func NewUserService(db *DB) *UserService {
	return &UserService{db}
}