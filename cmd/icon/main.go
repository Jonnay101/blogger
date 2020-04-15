package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/Jonnay101/icon/pkg/database"
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

	blog := blog.NewServer()
	blog.SetDatabase(database)

	srv := configureServer(getEnvOrDefault("PORT", "8080"), blog)
	log.Fatal(srv.ListenAndServe())

	return nil
}

func configureServer(port string, handler http.Handler) http.Server {

	return http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      handler,
	}
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
