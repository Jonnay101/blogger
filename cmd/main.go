package main

import (
	"net/http"
	"os"

	"github.com/Jonnay101/icon/pkg/blog"
	"github.com/Jonnay101/icon/pkg/database"
	"github.com/labstack/gommon/log"
)

func main() {
	// all packages will be initialized in here
	if err := Run(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

// Run -
func Run() error {
	// all normal main code in here with easy error returning
	database := database.NewDatabaseSession()
	blog := blog.NewServer(database)
	log.Fatal(http.ListenAndServe(":8080", blog))
	return nil
}
