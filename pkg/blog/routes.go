package blog

import "github.com/gorilla/mux"

func (s *server) setRoutes() {
	s.Router = mux.NewRouter()
	// handle creating new post
	s.Router.Handle("/blog", adminOnly(s.HandlerCreatePost())).Methods("POST")
	// handle getting single post request
	s.Router.Handle("/blog/{year}/{month}/{day}/{uuid}", s.HandlerGetPost()).Methods("GET")
	// handle getting multiple posts from corresponding partial urls
	s.Router.Handle("/blog/{year}/{month}/{day}/{uuid}", s.HandlerGetAllPosts()).Methods("GET")
	s.Router.Handle("/blog/{year}/{month}/{day}", s.HandlerGetAllPosts()).Methods("GET")
	s.Router.Handle("/blog/{year}/{month}", s.HandlerGetAllPosts()).Methods("GET")
	s.Router.Handle("/blog/{year}", s.HandlerGetAllPosts()).Methods("GET")
	s.Router.Handle("/blog", s.HandlerGetAllPosts()).Methods("GET")
}
