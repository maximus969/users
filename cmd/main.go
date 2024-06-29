package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/maximus969/users-app/internal/app/config"
	"github.com/maximus969/users-app/internal/app/repository/pgrepo"
	"github.com/maximus969/users-app/internal/app/services"
	"github.com/maximus969/users-app/internal/app/transport/grpcserver"
	"github.com/maximus969/users-app/internal/app/transport/httpserver"
	"github.com/maximus969/users-app/internal/pkg/pg"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	// read config from env
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
	userRepo := pgrepo.NewUserRepo(pgDB)
	userService := services.NewUserService(userRepo)

	// HTTP server setup
	httpServer := httpserver.NewHttpServer(userService)
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Users app"))
	}).Methods("GET")

	router.HandleFunc("/users", httpServer.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", httpServer.GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/user", httpServer.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", httpServer.UpdateUser).Methods(http.MethodPatch)
	router.HandleFunc("/user/{id}", httpServer.DeleteUser).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	grpcListener, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		return fmt.Errorf("failed to start grpc listener: %w", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpcserver.RegisterGRPCServer(grpcServer, &userService)

	errChan := make(chan error, 2)

	go func() {
		log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("HTTP server ListenAndServe Error: %v", err)
		}
	}()

	go func() {
		log.Printf("Starting gRPC server on %s", cfg.GRPCAddr)
		if err := grpcServer.Serve(grpcListener); err != nil {
			errChan <- fmt.Errorf("failed to serve gRPC server: %v", err)
		}
	}()

	// Handle graceful shutdown
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigint:
		log.Println("Shutting down servers...")
	case err := <-errChan:
		return err
	}

	// Gracefully stop HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("HTTP Server Shutdown Error: %v", err)
	}

	// Gracefully stop gRPC server
	grpcServer.GracefulStop()

	log.Println("Servers gracefully stopped")

	return nil
}

// runPgMigrations runs Postgres migrations
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

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
