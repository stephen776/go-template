package store

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

// DB is our struct that holds a database client
type DB struct {
	// Datasource Name
	DSN string

	client  *sqlx.DB
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

	db.client, err = sqlx.Connect("pgx", db.DSN)

	if err != nil {
		return err
	}

	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	db.cancel()

	// close database
	if db.client != nil {
		return db.client.Close()
	}

	return nil
}