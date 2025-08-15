package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mbarek-hani/rssagg/internal/database"
)

type apiConfig struct {
	Db *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		slog.Error("PORT is not found in the environment")
		os.Exit(1)
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		slog.Error("DB_URL is not found in the environment")
		os.Exit(1)
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	apiCfg := &apiConfig{
		Db: database.New(conn),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handleReadiness)
	v1router.Get("/error", handleError)
	v1router.Post("/users", apiCfg.handleCreateUser)
	v1router.Get("/users", apiCfg.middlwareAuth(apiCfg.handleGetUser))
	router.Mount("/v1", v1router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	slog.Info("the server is listening on http://localhost:" + port)

	err = server.ListenAndServe()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
