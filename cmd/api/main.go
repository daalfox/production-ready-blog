package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/daalfox/go-auth-api/pkg/auth"
	"github.com/daalfox/production-ready-blog/internal/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal("failed to connect to database")
	}

	authService := auth.NewAuthService(db)

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.Json)
	r.Mount("/users", authService.Router)

	fmt.Printf("server is running on port %v", ":8080")
	http.ListenAndServe(":8080", r)
}
