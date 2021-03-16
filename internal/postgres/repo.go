package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Repo is our struct that consturcts our Repo connection and creates sqlc Queries object
type Repo struct {
	// Datasource Name
	DSN string

	// sqlc generated Queries repo
	*Queries

	db *sql.DB
	ctx     context.Context // background context
	cancel  func()          // cancel background context
}

// Newr creates a new connection and returns the r object
func NewRepo(dsn string) *Repo {
	r := &Repo{ DSN: dsn}
	r.ctx, r.cancel = context.WithCancel(context.Background())

	return r
}

// Open opens the database connection with the DSN
func (r *Repo) Open() (err error) {
	if r.DSN == "" {
		return fmt.Errorf("dsn required to open r")
	}

	r.db, err = sql.Open("pgx", r.DSN)
	if err != nil {
		return err
	}

	r.Queries = New(r.db)

	return nil
}

// Close closes the database connection
func (r *Repo) Close() error {
	r.cancel()

	// close database
	if r.db != nil {
		return r.db.Close()
	}

	return nil
}