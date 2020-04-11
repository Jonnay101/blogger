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
	database, err := database.NewDatabaseSession()
	if err != nil {
		return err
	}
	blog := blog.NewServer()
	blog.SetDatabase(database)
	srv := configureServer(getPortFromEnvVars(), blog)
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

func getPortFromEnvVars() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
