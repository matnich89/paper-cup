package main

import (
	"log"
	"os"

	"papercup-test/internal/service/auth"
	"papercup-test/internal/service/video"

	"github.com/go-chi/chi/v5"

	"papercup-test/internal/db"
)

func main() {

	database, err := db.NewDatabase(os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	defer database.Client.Close()

	err = database.Client.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	migrationsSource := os.Getenv("MIGRATION_SOURCE")

	if migrationsSource == "" {
		log.Fatalln("migration source is not specified")
	}

	if err = database.MigrateDB(migrationsSource); err != nil {
		log.Fatalln(err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	domain := os.Getenv("DOMAIN")

	videoService := video.NewService(database)
	authService := auth.NewAuthService(database, domain, jwtSecret)

	router := chi.NewRouter()

	logger := log.Logger{}

	app := newApp(router, authService, videoService, &logger, domain, jwtSecret)

	log.Fatalln(app.serve())
}
