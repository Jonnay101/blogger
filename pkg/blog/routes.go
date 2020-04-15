package blog

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) setRoutes() {
	s.Router = mux.NewRouter()

	// handle creating new post
	s.Router.Handle("/blog", adminOnly(s.HandlerCreatePost())).Methods(http.MethodPost)

	// handle getting single post request
	s.Router.Handle("/blog/{year}/{month}/{day}/{uuid}", s.HandlerGetPost()).Methods(http.MethodGet)

	// handle getting multiple posts from corresponding partial urls
	s.Router.Handle("/blog/{year}/{month}/{day}/{uuid}", s.HandlerGetAllPosts()).Methods(http.MethodGet)
	s.Router.Handle("/blog/{year}/{month}/{day}", s.HandlerGetAllPosts()).Methods(http.MethodGet)
	s.Router.Handle("/blog/{year}/{month}", s.HandlerGetAllPosts()).Methods(http.MethodGet)
	s.Router.Handle("/blog/{year}", s.HandlerGetAllPosts()).Methods(http.MethodGet)
	s.Router.Handle("/blog", s.HandlerGetAllPosts()).Methods(http.MethodGet)

	// handle updatind a blog post
	s.Router.Handle("/blog/{year}/{month}/{day}/{uuid}", s.HandlerUpdatePost()).Methods(http.MethodPut)

	// handle removing blog posts
	s.Router.Handle("/blog/{year}/{month}/{day}/{uuid}", s.HandlerDeletePost()).Methods(http.MethodDelete)
}

// TODO: setup dynamic routes properly with regex
