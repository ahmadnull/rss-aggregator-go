package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahmadnull/rss-aggregator-go/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not found.")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", internal.HandlerReadiness)
	v1Router.Get("/error", internal.HandlerError)

	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port: %v", portString)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
