package blog

import "github.com/gorilla/mux"

func (s *server) setRoutes() {
	s.Router = mux.NewRouter()
	s.Router.Handle("/blog/", adminOnly(s.HandlerCreatePost())).Methods("POST")
	s.Router.Handle("/blog/{year}/{month}/{day}/{uuid}", s.HandlerGetPost()).Methods("GET")
}
