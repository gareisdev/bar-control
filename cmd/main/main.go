package main

import (
	"log"
	"os"

	"github.com/gareisdev/bar-control/internal/adapters/repositories"
	"github.com/gareisdev/bar-control/internal/core/usecases"
	"github.com/gareisdev/bar-control/internal/platform/database"
	"github.com/gareisdev/bar-control/internal/platform/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Using system environment variables.")
	}

	cfg := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     5432,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	}

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}

	defer db.Close()

	menuRepo := repositories.NewMenuRepository(db)
	menuUsecase := usecases.NewMenuUsecase(menuRepo)

	e := http.SetupServer(menuUsecase)

	log.Println("Application is running")
	log.Fatal(e.Start(":8080"))
}
