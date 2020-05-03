package main

import (
	"net/http"

	"github.com/Jonnay101/icon/pkg/auth"
	"github.com/Jonnay101/icon/pkg/handlers"
	"github.com/gorilla/mux"
)

func (srv *serverHandler) SetRoutes(iconHandlers *handlers.Handlers) {
	srv.Router = mux.NewRouter()

	// handle creating new post
	srv.Router.Handle("/blog/{user_uuid}", adminOnly(iconHandlers.HandlerBlogCreatePost())).Methods(http.MethodPost)

	// handle getting single post request
	srv.Router.Handle("/blog/{user_uuid}/{year}/{month}/{day}/{uuid}", iconHandlers.HandlerBlogGetPost()).Methods(http.MethodGet)

	// handle getting multiple posts from corresponding partial urls
	srv.Router.Handle("/blog/{user_uuid}/{year}/{month}/{day}/{uuid}", iconHandlers.HandlerBlogGetAllPosts()).Methods(http.MethodGet)
	srv.Router.Handle("/blog/{user_uuid}/{year}/{month}/{day}", iconHandlers.HandlerBlogGetAllPosts()).Methods(http.MethodGet)
	srv.Router.Handle("/blog/{user_uuid}/{year}/{month}", iconHandlers.HandlerBlogGetAllPosts()).Methods(http.MethodGet)
	srv.Router.Handle("/blog/{user_uuid}/{year}", iconHandlers.HandlerBlogGetAllPosts()).Methods(http.MethodGet)
	srv.Router.Handle("/blog/{user_uuid}", iconHandlers.HandlerBlogGetAllPosts()).Methods(http.MethodGet)

	// handle updatind a blog post
	srv.Router.Handle("/blog/{user_uuid}/{year}/{month}/{day}/{uuid}", iconHandlers.HandlerBlogUpdatePost()).Methods(http.MethodPut)

	// handle removing blog posts
	srv.Router.Handle("/blog/{user_uuid}/{year}/{month}/{day}/{uuid}", iconHandlers.HandlerBlogDeletePost()).Methods(http.MethodDelete)
}

func adminOnly(originalHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !testUserIsAdmin(true) {
			http.NotFound(w, r)
			return
		}

		originalHandler(w, r)
	}
}

func testUserIsAdmin(b bool) bool {

	user := &auth.User{}
	user.SetUserIsAdmin(true)

	return user.IsAdmin()
}
