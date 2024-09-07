package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cronnoss/tickets-api/internal/app/config"
	"github.com/cronnoss/tickets-api/internal/app/repository/pgrepo"
	"github.com/cronnoss/tickets-api/internal/app/services"
	"github.com/cronnoss/tickets-api/internal/app/transport/httpserver"
	"github.com/cronnoss/tickets-api/internal/pkg/pg"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	cfg := config.Read()

	pgDB, err := pg.Dial(cfg.DSN)
	if err != nil {
		return fmt.Errorf("pg.Dial failed: %w", err)
	}

	// run Postgres migrations
	if pgDB != nil {
		log.Println("Running PostgreSQL migrations")
		if err := runPgMigrations(cfg.DSN, cfg.MigrationsPath); err != nil {
			return fmt.Errorf("runPgMigrations failed: %w", err)
		}
	}

	// create repositories
	showRepo := pgrepo.NewShowRepo(pgDB)
	eventRepo := pgrepo.NewEventRepo(pgDB)
	placeRepo := pgrepo.NewPlaceRepo(pgDB)

	showService := services.NewShowService(showRepo)
	eventService := services.NewEventService(eventRepo)
	placeService := services.NewPlaceService(placeRepo)

	// create http server with application injected
	httpServer := httpserver.NewHTTPServer(showService, eventService, placeService)

	// create http router
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Tickets API v0.1"))
	}).Methods("GET")

	router.HandleFunc("/shows", httpServer.GetShows).Methods(http.MethodGet)
	router.HandleFunc("/shows/{id:[0-9]+}/events", httpServer.GetEvents).Methods(http.MethodGet)

	router.HandleFunc("/events/{id:[0-9]+}/places", httpServer.GetPlaces).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("Have a nice day!")

	return nil
}

// runPgMigrations runs Postgres migrations.
func runPgMigrations(dsn, path string) error {
	if path == "" {
		return errors.New("no migrations path provided")
	}
	if dsn == "" {
		return errors.New("no DSN provided")
	}

	m, err := migrate.New(
		path,
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
