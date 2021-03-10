package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"bitbucket.org/truefit/tf-manifest/internal/config"
	"bitbucket.org/truefit/tf-manifest/internal/http"
	"bitbucket.org/truefit/tf-manifest/internal/store"
	"bitbucket.org/truefit/tf-manifest/pkg/models"
)

// Main represents our program
type Main struct {
	Config *config.Config
	DB     *store.DB
	Server *http.Server

	// TODO: logger?

	// Services
	UserService models.UserService
}

func main() {
	// load the app config
	config, err := config.Load(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Setup signal handlers.
	// TODO: what else needs wired up for shutdown?
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
		DB:     store.NewDB(dsn),
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

	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			return err
		}
	}

	return nil
}

// Run starts the DB connection, injects services, runs http server
func (m *Main) Run(ctx context.Context) (err error) {
	if err := m.DB.Open(); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	// server config

	// inject services

	// start HTTP server
	if err := m.Server.Open(); err != nil {
		return err
	}

	log.Printf("running! dsn=%q", m.DB.DSN)

	return nil
}