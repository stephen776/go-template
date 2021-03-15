package store

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB is our struct that consturcts our DB connection and creates sqlc Queries object
type DB struct {
	// Datasource Name
	DSN string

	// sqlc generated Queries repo
	Queries *Queries

	db *sql.DB
	ctx     context.Context // background context
	cancel  func()          // cancel background context
}

// NewDB creates a new connection and returns the DB object
func NewDB(dsn string) *DB {
	db := &DB{ DSN: dsn}
	db.ctx, db.cancel = context.WithCancel(context.Background())

	return db
}

// Open opens the database connection with the DSN
func (db *DB) Open() (err error) {
	if db.DSN == "" {
		return fmt.Errorf("dsn required to open db")
	}

	db.db, err = sql.Open("pgx", db.DSN)
	if err != nil {
		return err
	}

	db.Queries = New(db.db)

	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	db.cancel()

	// close database
	if db.db != nil {
		return db.db.Close()
	}

	return nil
}