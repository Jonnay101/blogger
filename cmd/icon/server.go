package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jonnay101/icon/pkg/handlers"
	"github.com/gorilla/mux"
)

type serverHandler struct {
	Router *mux.Router
}

// NewServer -
func NewServer(port string, iconHandlers handlers.Handlers) *http.Server {

	server := serverHandler{}

	server.SetRoutes(iconHandlers)

	httpServer := server.configureServer(port)

	return &httpServer
}

func (srv *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.Router.ServeHTTP(w, r)
}

func (srv *serverHandler) configureServer(port string) http.Server {

	return http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      srv,
	}
}
