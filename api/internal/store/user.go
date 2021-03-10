package store

import (
	"github.com/jmoiron/sqlx"

	"bitbucket.org/truefit/tf-manifest/pkg/models"
)

// UserService performs DB operations against the users table
type UserService struct {
	db *sqlx.DB
}

// GetUsers returns all users
func (s *UserService) GetUsers() ([]models.User, error) {
	users := []models.User{}

	err := s.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// NewUserService creates a new instance of the UserRepository
func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{db}
}