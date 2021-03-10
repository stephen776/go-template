package models

// User represents a User in our domain
type User struct {
	ID        int    `db:"id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}


// UserService defines CRUD for User model
type UserService interface {
	GetUsers() ([]*User, error)
	// GetUserById(id int) (*User, error)
	// CreateUser(u *User) error
	// DeleteUser(id int) error
}