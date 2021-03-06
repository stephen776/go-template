package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/stephen776/go-template/internal/config"
	"github.com/stephen776/go-template/internal/domain"
	"github.com/stephen776/go-template/internal/http"
	"github.com/stephen776/go-template/internal/postgres"
)

// TODO:
// - run migrations on start
// - Auth
// - request validation

// Main represents our program
type Main struct {
	Config *config.Config
	Repo     *postgres.Repo
	Server *http.Server

	// TODO: logger?

	// Services
	UserService domain.UserService
}

func main() {
	// load the app config
	config, err := config.Load(".")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	m := NewMain(config)

	// Execute program.
	if err := m.Run(ctx); err != nil {
		m.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Wait for CTRL-C.
	<-ctx.Done()

	// clean up
	if err := m.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// NewMain returns a new instance of our app
func NewMain(config *config.Config) *Main {
	// TODO: maybe we move this?
	// init the database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable",
			config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)

	return &Main{
		Config: config,
		Repo:     postgres.NewRepo(dsn),
		Server: http.NewServer(),
	}
}

// Close gracefully stops the program.
func (m *Main) Close() error {
	if m.Server != nil {
		if err := m.Server.Close(); err != nil {
			return err
		}
	}

	if m.Repo != nil {
		if err := m.Repo.Close(); err != nil {
			return err
		}
	}

	return nil
}

// Run starts the DB connection, injects services, runs http server
func (m *Main) Run(ctx context.Context) (err error) {
	if err := m.Repo.Open(); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	// server config

	// instantiate services

	// inject services
	m.Server.UserService = m.Repo

	// start HTTP server
	m.Server.Addr = m.Config.ServerPort

	if err := m.Server.Open(); err != nil {
		return err
	}

	log.Printf("running! dsn=%q", m.Repo.DSN)

	return nil
}