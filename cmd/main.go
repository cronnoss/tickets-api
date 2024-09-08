package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/cronnoss/tickets-api/docs"
	"github.com/cronnoss/tickets-api/internal/app/config"
	"github.com/cronnoss/tickets-api/internal/app/repository/memory"
	"github.com/cronnoss/tickets-api/internal/app/services"
	"github.com/cronnoss/tickets-api/internal/app/transport/httpserver"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

// @title Tickets API
// @version 0.1
// @description API Server for remote Tickets Application
func run() error { // nolint unparam
	cfg := config.Read()

	// create repositories
	showRepo := memory.NewShowRepo()
	eventRepo := memory.NewEventRepo()
	placeRepo := memory.NewPlaceRepo()

	showService := services.NewShowService(&showRepo)
	eventService := services.NewEventService(&eventRepo)
	placeService := services.NewPlaceService(&placeRepo)

	// create http server with application injected
	httpServer := httpserver.NewHTTPServer(showService, eventService, placeService)

	// create http router
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Tickets API v0.1"))
	}).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

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
