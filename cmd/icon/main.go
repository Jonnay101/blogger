package main

import (
	"fmt"
	"os"

	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/Jonnay101/icon/pkg/comments"
	"github.com/Jonnay101/icon/pkg/database"
	"github.com/Jonnay101/icon/pkg/handlers"
	"github.com/labstack/gommon/log"
)

func main() {

	if err := Run(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

// Run -
func Run() error {

	mongoURL, err := mustenv("MONGO_URL")
	if err != nil {
		return err
	}

	database, err := database.NewDatabaseSession(mongoURL)
	if err != nil {
		return err
	}

	blogService := blog.NewService(database)

	commentsService := comments.NewService(database)

	iconHandlers := handlers.NewHandlers(blogService, commentsService)

	server := NewServer(getEnvOrDefault("PORT", "8080"), iconHandlers)

	log.Fatal(server.ListenAndServe())

	return nil
}

func getEnvOrDefault(envVar, def string) string {

	ev := os.Getenv(envVar)
	if envVar == "" {
		ev = def
	}

	return ev
}

func mustenv(envVar string) (string, error) {

	ev := os.Getenv(envVar)
	if ev == "" {
		return "", fmt.Errorf("Cannot find %s environment variable", envVar)
	}

	return ev, nil
}
