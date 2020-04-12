package blog

import "github.com/gorilla/mux"

func (s *server) setRoutes() {
	s.Router = mux.NewRouter()
	s.Router.Handle("/", s.HandlerCreatePost()).Methods("POST")
}
