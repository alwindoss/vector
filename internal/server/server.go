package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alwindoss/vector/internal/station"
	"github.com/alwindoss/vector/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() error {
	err := godotenv.Load()
	if err != nil {
		err = fmt.Errorf("unable to load environment vairables: %w", err)
		return err
	}
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUserName := os.Getenv("DB_UNAME")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUserName, dbPasswd, dbName, dbPort, dbSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err = fmt.Errorf("unable to connect to DB: %w", err)
		return err
	}
	storage.Setup(db)
	repo := storage.NewStationRepository(db)
	svc := station.NewService(repo)
	h := NewHandler(svc)
	setupRoutes(r, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	addr := fmt.Sprintf(":%s", port)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		err = fmt.Errorf("unable to ListenAndServe on port %s: %w", port, err)
		return err
	}
	return nil
}

func setupRoutes(r chi.Router, h Handler) {
	r.Route("/vector/api/v1", func(r chi.Router) {
		r.Post("/stations", h.CreateStation)
	})
}
